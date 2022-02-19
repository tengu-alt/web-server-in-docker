package service

import (
	"errors"
	"registration-web-service2/pkg/models"
	"registration-web-service2/pkg/store"
	"registration-web-service2/pkg/validation"
)

func SignUp(u models.User) ([]models.ValidationErr, error) {
	var err error
	validErrors := validation.Validate(u)
	if len(validErrors) > 0 {
		err = errors.New("failedValidation")
		return validErrors, err

	}
	salt, hash := ReturnSaltAndHash(u.Password)
	store.InsertToDB(u, salt, hash)
	return nil, nil

}

//func Login(u models.LoginUser) (models.TokenResponse,error){
//
//}
