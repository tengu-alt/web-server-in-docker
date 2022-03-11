package validation

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"time"
	"web-server-in-docker/pkg/models"
	"web-server-in-docker/pkg/store"
)

type User = models.User
type LoginUser = models.LoginUser
type ValidationErr = models.ValidationErr

func ValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(email) != true {
		return false
	}
	return true
}

func CheckMail(email string) bool {
	pgxconn, err := pgx.Connect(context.Background(), store.GetConfig())
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

func LoginValid(u LoginUser) bool {
	if CheckMail(u.LoginMail) == false {
		if CheckLoginPassword(u) == true {
			return true
		}
	}
	return false
}

func CheckLoginPassword(u LoginUser) bool {
	pgxconn, err := pgx.Connect(context.Background(), store.GetConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pgxconn.Close(context.Background())
	var searchId int
	err = pgxconn.QueryRow(context.Background(), "SELECT user_id FROM signed_users WHERE email=$1", u.LoginMail).Scan(&searchId)
	var searchSalt, searchHash string
	err = pgxconn.QueryRow(context.Background(), "SELECT salt, salt_hash FROM credentials WHERE user_id=$1", searchId).Scan(&searchSalt, &searchHash)
	compare, _ := base64.StdEncoding.DecodeString(searchHash)
	err = bcrypt.CompareHashAndPassword(compare, []byte(u.LoginPassword+searchSalt))
	if err != nil {

		return false
	}

	return true
}

func TimeValid(i interface{}) bool {
	lifeTime := i.(float64)

	if int64(lifeTime) < time.Now().Unix() {
		fmt.Println("time!")
		return false
	}
	return true

}
