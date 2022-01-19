package store

import (
	"database/sql"
	"fmt"
	"registration-web-service2/pkg/users"

	_ "github.com/lib/pq"
)

type User = users.User

const (
	host     = "localhost"
	port     = 6080
	user     = "postgres"
	password = "12345"
	dbname   = "users"
)

func InsertToDB(u User) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO signed_users (firstname,lastname,email,user_password) VALUES($1,$2,$3,$4)", u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	fmt.Println("inserting")
}
