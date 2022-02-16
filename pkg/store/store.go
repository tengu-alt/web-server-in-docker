package store

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"registration-web-service2/pkg/users"
	"time"
)

type User = users.User
type Config = users.Config
type LoginUser = users.LoginUser

func GetConfig() string {
	yfile, err := ioutil.ReadFile("../cmd/config.yaml")

	if err != nil {

		log.Fatal(err)
	}

	conf := *&users.Config{}

	err2 := yaml.Unmarshal(yfile, &conf)

	if err2 != nil {

		log.Fatal(err2)
	}
	result := fmt.Sprintf("postgres://%s:%s@sqlserver/%s?sslmode=disable", conf.User, conf.Password, conf.DB)
	return result
}

func GetKey() string {
	yfile, err := ioutil.ReadFile("../cmd/config.yaml")

	if err != nil {

		log.Fatal(err)
	}

	conf := *&users.Config{}

	err2 := yaml.Unmarshal(yfile, &conf)

	if err2 != nil {

		log.Fatal(err2)
	}
	return conf.Key
}

func InsertToDB(u User) {
	db, err := sqlx.Connect("postgres", GetConfig())
	if err != nil {
		log.Fatalln(err)
	}
	//pgxconn, err := pgx.Connect(context.Background(), GetConfig())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//psqlconn := fmt.Sprintf(GetConfig())
	//db, err := sql.Open("postgres", psqlconn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	_, err = db.Queryx("INSERT INTO signed_users (firstname,lastname,email) VALUES($1,$2,$3)", u.FirstName, u.LastName, u.Email)
	if err != nil {
		panic(err)
	}
	//defer insert.Close()
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	var searchId int
	err = db.Get(&searchId, "SELECT user_id FROM signed_users WHERE email=$1", u.Email)
	_, err = db.Queryx("insert into credentials (user_id,salt,salt_hash) values ($1,$2,$3)", searchId, base64.StdEncoding.EncodeToString(salt), HashPassword(u.Password+base64.StdEncoding.EncodeToString(salt)))
	if err != nil {
		panic(err)
	}
	//defer insertHash.Close()
	fmt.Println("inserting")

}

func HashPassword(password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(hashedPassword)

}

func GiveToken(u LoginUser) string {
	db, err := sqlx.Connect("postgres", GetConfig())
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

func DropToken(token string) {
	//fmt.Println(token)
	//psqlconn := fmt.Sprintf(GetConfig())
	//db, err := sql.Open("postgres", psqlconn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	db, err := sqlx.Connect("postgres", GetConfig())
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Queryx("delete from tokens where token = $1", token)
	if err != nil {
		panic(err)
	}
	//defer drop.Close()
}

func CreateToken(email string) string {
	//pgxconn, err := pgx.Connect(context.Background(), GetConfig())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	db, err := sqlx.Connect("postgres", GetConfig())
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
	hmacSampleSecret := []byte(GetKey())
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

func TestDb() string {
	db, err := sqlx.Connect("postgres", GetConfig())
	if err != nil {
		log.Fatalln(err)
	}
	var pers string
	err = db.Get(&pers, "SELECT lastname FROM signed_users WHERE user_id=1")
	if err == nil {
		fmt.Sprintln(err)
	}
	return pers
}
