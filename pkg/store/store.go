package store

import (
	"database/sql"
	"fmt"
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
	fmt.Println("inserting")
}
