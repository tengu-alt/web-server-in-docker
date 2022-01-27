package store

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"registration-web-service2/pkg/users"

	_ "github.com/lib/pq"
)

type User = users.User

func InsertToDB(u User) {
	pgxconn, err := pgx.Connect(context.Background(), "postgres://postgres:12345@sqlserver/users?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
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
	var searchId int
	err = pgxconn.QueryRow(context.Background(), "SELECT user_id FROM signed_users WHERE email=$1", u.Email).Scan(&searchId)
	insertHash, err := db.Query("insert into credentials (user_id,salt,salt_hash) values ($1,$2,$3)", searchId, base64.StdEncoding.EncodeToString(salt), HashPassword(u.Password+base64.StdEncoding.EncodeToString(salt)))
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
