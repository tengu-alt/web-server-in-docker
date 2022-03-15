package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"web-server-in-docker/pkg/models"
	"web-server-in-docker/pkg/store"
	"web-server-in-docker/pkg/validation"
)

func SignUp(u models.User, conn *sqlx.DB) ([]models.ValidationErr, error) {
	var err error
	validErrors := validation.Validate(u)
	if len(validErrors) > 0 {
		err = errors.New("failedValidation")
		return validErrors, err

	}
	salt, hash := ReturnSaltAndHash(u.Password)
	err = store.InsertToDB(u, salt, hash, conn)
	if err != nil {
		panic(err)
	}
	return nil, nil

}

func Login(u models.LoginUser, conn *sqlx.DB) (models.TokenResponse, error) {
	var err error
	resp := models.TokenResponse{}
	if validation.LoginValid(u) == true {
		token := GiveToken(u, conn)
		resp.ResponseMessage = "success login"
		resp.Token = token
		return resp, nil
	} else {
		resp.ResponseMessage = "invalid data"
		err = errors.New("failedValidation")
		return resp, err
	}
}
func Logout(token string, conn *sqlx.DB) error {
	err := store.DropToken(token, conn)
	if err != nil {
		return err
	}
	return nil
}

func SayName(tokenString string) (string, bool) {
	hmacSampleSecret := []byte(store.GetKey())
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	var result string

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if validation.TimeValid(claims["ExpiresAt"]) {
			result = fmt.Sprintln(claims["name"])
		} else {
			return "", false
		}
	} else {
		fmt.Println(err)
		return "", false
	}
	return result, true
}
