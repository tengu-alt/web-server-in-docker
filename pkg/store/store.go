package store

//"postgres://postgres:12345@localhost/models?sslmode=disable"
//migrate -database postgres://postgres:12345@localhost/models?sslmode=disable -path . up
import (
	"encoding/base64"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"web-server-in-docker/pkg/models"
)

type User = models.User

type LoginUser = models.LoginUser

type DataBase struct {
	DBmodel *sqlx.DB
}

func NewConnect(str string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", str)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (conn *DataBase) InsertToDB(u User, salt, hash string) error {
	db := conn.DBmodel

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

func (conn *DataBase) DropToken(token string) error {

	db := conn.DBmodel
	var searchToken string
	err := db.Get(&searchToken, "select token from tokens where token = $1", token)
	if err != nil {
		return err
	}
	if token != searchToken {
		return err

	}
	_, err = db.Queryx("delete from tokens where token = $1", token)
	if err != nil {
		return err
	}
	return nil
}

func (conn *DataBase) GetNames(email string) (string, string, error) {

	db := conn.DBmodel
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

func (conn *DataBase) InsertToken(email, token string) error {
	db := conn.DBmodel
	var searchId int
	err := db.Get(&searchId, "SELECT user_id FROM signed_users WHERE email=$1", email)
	if err != nil {
		return err
	}
	_, err = db.Queryx("insert into tokens (user_id,token) values ($1,$2)", searchId, token)
	if err != nil {
		return err
	}
	return nil

}

func (conn *DataBase) CheckMail(email string) bool {

	db := conn.DBmodel
	var searchMail string
	err := db.Get(&searchMail, "SELECT email FROM signed_users WHERE email=$1", email)
	if err != nil {
		fmt.Println(err)
	}
	if searchMail == email {
		return false
	}
	return true
}

func (conn *DataBase) CheckLoginPassword(u LoginUser) bool {
	db := conn.DBmodel
	var searchId int
	err := db.Get(&searchId, "SELECT user_id FROM signed_users WHERE email=$1", u.LoginMail)
	var searchSalt, searchHash string
	err = db.Get(&searchSalt, "SELECT salt FROM credentials WHERE user_id=$1", searchId)
	err = db.Get(&searchHash, "SELECT salt_hash FROM credentials WHERE user_id=$1", searchId)
	compare, _ := base64.StdEncoding.DecodeString(searchHash)
	err = bcrypt.CompareHashAndPassword(compare, []byte(u.LoginPassword+searchSalt))
	if err != nil {

		return false
	}

	return true
}
