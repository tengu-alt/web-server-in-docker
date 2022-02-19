package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"log"
	"registration-web-service2/pkg/models"
	"registration-web-service2/pkg/store"
	"time"
)

func GiveToken(u models.LoginUser) string {
	db, err := sqlx.Connect("postgres", store.GetConfig())
	if err != nil {
		log.Fatalln(err)
	}
	//pgxconn, err := pgx.Connect(context.Background(), GetConfig())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer pgxconn.Close(context.Background())
	//psqlconn := fmt.Sprintf(GetConfig())
	//db, err := sql.Open("postgres", psqlconn)
	//if err != nil {
	//panic(err)
	//}
	//defer db.Close()
	var searchId int
	err = db.Get(&searchId, "SELECT user_id FROM signed_users WHERE email=$1", u.LoginMail)
	token := CreateToken(u.LoginMail)
	_, err = db.Queryx("insert into tokens (user_id,token) values ($1,$2)", searchId, token)
	if err != nil {
		panic(err)
	}
	//defer insertToken.Close()
	fmt.Println("inserting token")
	return token
}

func CreateToken(email string) string {
	//pgxconn, err := pgx.Connect(context.Background(), GetConfig())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	db, err := sqlx.Connect("postgres", store.GetConfig())
	if err != nil {
		log.Fatalln(err)
	}
	var Fname string
	var Lname string
	err = db.Get(&Fname, "SELECT firstname FROM signed_users WHERE email=$1", email)
	if err == nil {
		fmt.Sprintln(err)
	}
	err = db.Get(&Lname, "SELECT lastname FROM signed_users WHERE email=$1", email)
	if err == nil {
		fmt.Sprintln(err)
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
