package validation

import (
	"fmt"
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
	}

	return errors
}
