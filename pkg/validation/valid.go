package validation

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"regexp"
	"registration-web-service2/pkg/users"
)

type User = users.User
type ValidationErr = users.ValidationErr

func Printer(i string) string {
	fmt.Printf(i)
	return i

}
func ValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(email) != true {
		return false
	}
	return true
}
func CheckMail(email string) bool {
	pgxconn, err := pgx.Connect(context.Background(), "postgres://postgres:12345@sqlserver/users?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pgxconn.Close(context.Background())
	var searchMail string
	err = pgxconn.QueryRow(context.Background(), "SELECT email FROM signed_users WHERE email=$1", email).Scan(&searchMail)
	if searchMail == email {
		return false
	}
	return true
}
func NameValid(name string, min, max int) bool {
	if len(name) < min || len(name) > max {
		return false
	}
	return true
}
func PasswordValid(password string, min int) bool {
	if len(password) < min {
		return false
	}
	return true
}
func Validate(u User) []ValidationErr {
	errors := make([]ValidationErr, 0, 0)
	if NameValid(u.FirstName, 2, 64) != true {
		errors = append(errors, ValidationErr{
			FieldValue: "FirstName",
			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 2 and less than 64", "FirstName"),
		})
	}
	if NameValid(u.LastName, 2, 64) != true {
		errors = append(errors, ValidationErr{
			FieldValue: "LastName",
			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 2 and less than 64", "Lastname"),
		})
	}

	if PasswordValid(u.Password, 8) != true {
		errors = append(errors, ValidationErr{
			FieldValue: "Password",
			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 8", "Password"),
		})
	}

	if ValidEmail(u.Email) != true {
		errors = append(errors, ValidationErr{
			FieldValue: "Email",
			ErrMassage: "email failed verification",
		})
	} else if CheckMail(u.Email) != true {
		errors = append(errors, ValidationErr{
			FieldValue: "Email",
			ErrMassage: "email is already exist",
		})
	}

	return errors
}
