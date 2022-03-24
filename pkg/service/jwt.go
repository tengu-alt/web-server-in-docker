package service

import (
	"github.com/golang-jwt/jwt"
	"time"
	"web-server-in-docker/pkg/configs"
	"web-server-in-docker/pkg/models"
	"web-server-in-docker/pkg/store"
)

func GiveToken(u models.LoginUser, conn *store.DataBase) string {
	token := CreateToken(u.LoginMail, conn)
	err := conn.InsertToken(u.LoginMail, token)
	if err != nil {
		return err.Error()
	}
	return token
}

func CreateToken(email string, conn *store.DataBase) string {
	Fname, Lname, err := conn.GetNames(email)
	if err != nil {
		return err.Error()
	}
	hmacSampleSecret := []byte(configs.GetKey())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      Fname + " " + Lname,
		"ExpiresAt": time.Now().Add(12 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return tokenString
}
