package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"time"
	"web-server-in-docker/pkg/models"
	"web-server-in-docker/pkg/store"
)

func GiveToken(u models.LoginUser, conn *sqlx.DB) string {
	token := CreateToken(u.LoginMail, conn)
	err := store.InsertToken(u.LoginMail, token, conn)
	if err != nil {
		return err.Error()
	}
	return token
}

func CreateToken(email string, conn *sqlx.DB) string {
	Fname, Lname, err := store.GetNames(email, conn)
	if err != nil {
		return err.Error()
	}
	hmacSampleSecret := []byte(store.GetKey())
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
