package store

//"postgres://postgres:12345@localhost/models?sslmode=disable"
//migrate -database postgres://postgres:12345@localhost/models?sslmode=disable -path . up
import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"web-server-in-docker/pkg/models"
)

type User = models.User
type Config = models.Config
type LoginUser = models.LoginUser

func NewConnect(str string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", str)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetConfig() string {
	yfile, err := ioutil.ReadFile("../cmd/config.yaml")

	if err != nil {

		log.Fatal(err)
	}
	conf := &Config{}

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
	conf := &Config{}
	err2 := yaml.Unmarshal(yfile, &conf)
	if err2 != nil {

		log.Fatal(err2)
	}
	return conf.Key
}

func InsertToDB(u User, salt, hash string, conn *sqlx.DB) error {
	db := conn

	_, err := db.Queryx("INSERT INTO signed_users (firstname,lastname,email) VALUES($1,$2,$3)", u.FirstName, u.LastName, u.Email)
	if err != nil {
		return err
	}

	var searchId int
	err = db.Get(&searchId, "SELECT user_id FROM signed_users WHERE email=$1", u.Email)
	_, err = db.Queryx("insert into credentials (user_id,salt,salt_hash) values ($1,$2,$3)", searchId, salt, hash)
	if err != nil {
		return err
	}
	return nil
}

func DropToken(token string, conn *sqlx.DB) error {

	db := conn
	_, err := db.Queryx("delete from tokens where token = $1", token)
	if err != nil {
		return err
	}
	return nil
}

func GetNames(email string, conn *sqlx.DB) (string, string, error) {

	db := conn
	var Fname string
	var Lname string
	err := db.Get(&Fname, "SELECT firstname FROM signed_users WHERE email=$1", email)
	if err != nil {
		return "", "", err
	}
	err = db.Get(&Lname, "SELECT lastname FROM signed_users WHERE email=$1", email)
	if err != nil {
		return "", "", err
	}
	return Fname, Lname, nil
}

func InsertToken(email, token string, conn *sqlx.DB) error {
	db := conn
	var searchId int
	err := db.Get(&searchId, "SELECT user_id FROM signed_users WHERE email=$1", email)
	_, err = db.Queryx("insert into tokens (user_id,token) values ($1,$2)", searchId, token)
	if err != nil {
		return err
	}
	return nil

}
