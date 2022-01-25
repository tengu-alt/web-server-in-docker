package store

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"registration-web-service2/pkg/users"

	_ "github.com/lib/pq"
)

type User = users.User

func InsertToDB(u User) {
	psqlconn := fmt.Sprintf("postgres://postgres:12345@sqlserver/users?sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO signed_users (firstname,lastname,email) VALUES($1,$2,$3)", u.FirstName, u.LastName, u.Email)
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	insertHash, err := db.Query("insert into credentials (salt,satl_hash) values ($1,$2)", base64.StdEncoding.EncodeToString(salt), HashPassword(u.Password+base64.StdEncoding.EncodeToString(salt)))
	if err != nil {
		panic(err)
	}
	defer insertHash.Close()
	fmt.Println("inserting")

}
func HashPassword(password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(hashedPassword)

}
