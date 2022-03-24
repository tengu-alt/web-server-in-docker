package validation

import (
	"fmt"
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

func Validate(u User, conn *store.DataBase) []ValidationErr {
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
	} else if conn.CheckMail(u.Email) != true {
		errors = append(errors, ValidationErr{
			FieldValue: "Email",
			ErrMassage: "email is already exist",
		})
	}

	return errors
}

func LoginValid(u LoginUser, conn *store.DataBase) bool {
	if conn.CheckMail(u.LoginMail) == false {
		if conn.CheckLoginPassword(u) == true {
			return true
		}
	}
	return false
}

func TimeValid(i interface{}) bool {
	lifeTime := i.(float64)

	if int64(lifeTime) < time.Now().Unix() {
		fmt.Println("time!")
		return false
	}
	return true

}
