package store

import (
	"fmt"
	"testing"
	"web-server-in-docker/pkg/models"
)

func TestInsertToDB(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
	}
	salt := "QFVgDYtUIzM="
	hash := "JDJhJDEwJDBvdy9rNk1RdGkwemFaZHYwbWhCRXVKU2x5NEd6ZnBlVGRBYUdsLjFXbWFIcVJ6d2EzNHQu"
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	err = InsertToDB(u, salt, hash, database)
	if err != nil {
		t.Error(err)
	}

}

func TestInsertToken(t *testing.T) {
	email := "123@123.123"
	token := "token"
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	err = InsertToken(email, token, database)
	if err != nil {
		t.Fail()
	}
}

func TestGetNames(t *testing.T) {
	email := "123@123.123"
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	fname, lname, err := GetNames(email, database)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fname, lname)
}
func TestDropToken(t *testing.T) {
	token := "token"
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	err = DropToken(token, database)
	if err != nil {
		t.Error(err)
	}

}
