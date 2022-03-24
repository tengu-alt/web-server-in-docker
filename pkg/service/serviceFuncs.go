package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"web-server-in-docker/pkg/configs"
	"web-server-in-docker/pkg/models"
	"web-server-in-docker/pkg/store"
	"web-server-in-docker/pkg/validation"
)

func SignUp(u models.User, conn *store.DataBase) ([]models.ValidationErr, error) {
	var err error
	validErrors := validation.Validate(u, conn)
	if len(validErrors) > 0 {
		err = errors.New("failedValidation")
		return validErrors, err

	}
	salt, hash := ReturnSaltAndHash(u.Password)
	err = conn.InsertToDB(u, salt, hash)
	if err != nil {
		panic(err)
	}
	return nil, nil

}

func Login(u models.LoginUser, conn *store.DataBase) (models.TokenResponse, error) {
	var err error
	resp := models.TokenResponse{}
	if validation.LoginValid(u, conn) == true {
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
func Logout(token string, conn *store.DataBase) error {
	err := conn.DropToken(token)
	if err != nil {
		return err
	}
	return nil
}

func SayName(tokenString string) (string, bool) {
	hmacSampleSecret := []byte(configs.GetKey())
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
