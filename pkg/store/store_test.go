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

type testCase struct {
	User  models.User
	salt  string
	hash  string
	token string
}

//func truncateCustomTable(name string) {
//	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
//	database, _ := NewConnect(connString)
//	conn := getConnStruct(database)
//	db := conn.DBmodel
//	_, err := db.Queryx("TRUNCATE TABLE $1", name)
//	if err != nil {
//		panic(err)
//	}
//}

func Test_InsertUserToPositive(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
	}
	testcaseDBPos := testCase{
		u,
		"QFVgDYtUIzM=",
		"JDJhJDEwJDBvdy9rNk1RdGkwemFaZHYwbWhCRXVKU2x5NEd6ZnBlVGRBYUdsLjFXbWFIcVJ6d2EzNHQu",
		"asdasadasd",
	}
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	conn := getConnStruct(database)
	err = conn.InsertToDB(testcaseDBPos.User, testcaseDBPos.salt, testcaseDBPos.hash)
	if err != nil {
		t.Fatal("cannot insert")
	}
}

func TestDataBase_InsertToken(t *testing.T) {
	u := models.User{
		Email: "123@123.1231",
	}
	testestcaseDBPos := testCase{
		User:  u,
		token: "asdasadasd",
	}
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	conn := getConnStruct(database)
	err = conn.InsertToken(testestcaseDBPos.User.Email, testestcaseDBPos.token)
	if err != nil {
		t.Fatal("cannot insert")
	}
}
func TestDataBase_DropToken(t *testing.T) {
	testestcaseDBPos := testCase{
		token: "asdasadasd",
	}
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	conn := getConnStruct(database)
	err = conn.DropToken(testestcaseDBPos.token)
	if err != nil {
		t.Fatal("cannot insert")
	}
}
func TestDataBase_CheckMail(t *testing.T) {
	u := models.User{
		Email: "123@123.1231",
	}
	testestcaseDBPos := testCase{
		User: u,
	}
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	conn := getConnStruct(database)
	err2 := conn.CheckMail(testestcaseDBPos.User.Email)
	if err2 {
		t.Fatal("cannot check")
	}
}
func TestDataBase_GetNames(t *testing.T) {
	u := models.User{
		Email: "123@123.1231",
	}
	testestcaseDBPos := testCase{
		User: u,
	}
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, err := NewConnect(connString)
	if err != nil {
		t.Error(err)
	}
	conn := getConnStruct(database)
	_, _, err = conn.GetNames(testestcaseDBPos.User.Email)
	if err != nil {
		t.Fatal("cannot get names")
	}
}
