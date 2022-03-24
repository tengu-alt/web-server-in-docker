package store

import (
	"github.com/jmoiron/sqlx"
	"testing"
	"web-server-in-docker/pkg/models"
)

func getConnStruct(db *sqlx.DB) *DataBase {
	return &DataBase{
		db,
	}
}

func TestDBPositive(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
	}
	salt := "QFVgDYtUIzM="
	hash := "JDJhJDEwJDBvdy9rNk1RdGkwemFaZHYwbWhCRXVKU2x5NEd6ZnBlVGRBYUdsLjFXbWFIcVJ6d2EzNHQu"
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	token := "asdasadasd"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	conn := getConnStruct(database)
	t.Run("insert", func(t *testing.T) {
		err := conn.InsertToDB(u, salt, hash)
		if err != nil {
			t.Fatal("cannot insert")
		}
	})
	t.Run("insertToken", func(t *testing.T) {
		err := conn.InsertToken(u.Email, token)
		if err != nil {
			t.Fatal("cannot insert")
		}
	})
	t.Run("dropToken", func(t *testing.T) {
		err := conn.DropToken(token)
		if err != nil {
			t.Fatal("cannot drop")
		}
	})
	t.Run("getNames", func(t *testing.T) {
		_, _, err := conn.GetNames(u.Email)
		if err != nil {
			t.Fatal("cannot get names")
		}
	})
	t.Run("dropToken", func(t *testing.T) {
		err := conn.DropToken(token)
		if err != nil {
			t.Fatal("cannot drop")
		}
	})
	t.Run("checkEmail", func(t *testing.T) {
		err := conn.CheckMail(u.Email)
		if err {
			t.Fatal("cannot check")
		}
	})

}
