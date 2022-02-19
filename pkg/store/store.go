package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"registration-web-service2/pkg/models"
)

type User = models.User
type Config = models.Config
type LoginUser = models.LoginUser

func GetConfig() string {
	yfile, err := ioutil.ReadFile("../cmd/config.yaml")

	if err != nil {

		log.Fatal(err)
	}

	conf := *&models.Config{}

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

	conf := *&models.Config{}

	err2 := yaml.Unmarshal(yfile, &conf)

	if err2 != nil {

		log.Fatal(err2)
	}
	return conf.Key
}

func InsertToDB(u User, salt, hash string) {
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
	//salt := make([]byte, 8)
	//if _, err := rand.Read(salt); err != nil {
	//panic(err)
	//}
	var searchId int
	err = db.Get(&searchId, "SELECT user_id FROM signed_users WHERE email=$1", u.Email)
	_, err = db.Queryx("insert into credentials (user_id,salt,salt_hash) values ($1,$2,$3)", searchId, salt, hash)
	if err != nil {
		panic(err)
	}
	//defer insertHash.Close()

}

//func HashPassword(password string) string {
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		panic(err)
//	}
//
//	return base64.StdEncoding.EncodeToString(hashedPassword)
//
//}

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
