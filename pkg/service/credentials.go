package service

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

func ReturnSaltAndHash(password string) (string, string) {
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	hashedPassword := HashPassword(password + base64.StdEncoding.EncodeToString(salt))
	return base64.StdEncoding.EncodeToString(salt), hashedPassword
}
func HashPassword(password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(hashedPassword)

}
