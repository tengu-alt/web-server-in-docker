package store

import (
	"fmt"
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

func truncateCustomTable(name string) {
	connString := "postgres://postgres:12345@localhost:6080/models?sslmode=disable"
	database, _ := NewConnect(connString)
	conn := getConnStruct(database)
	db := conn.DBmodel
	_, err := db.Queryx(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", name))
	if err != nil {
		panic(err)
	}
}

func Test_InsertUserPositive(t *testing.T) {
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
		t.Error(err)
	}
	truncateCustomTable("signed_users")
}

func TestDataBase_InsertToken(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	err = conn.InsertToDB(testestcaseDBPos.User, "", "")
	err = conn.InsertToken(testestcaseDBPos.User.Email, testestcaseDBPos.token)
	if err != nil {
		t.Error(err)
	}
	truncateCustomTable("signed_users")
}
func TestDataBase_DropToken(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	err = conn.InsertToDB(testestcaseDBPos.User, "", "")
	err = conn.InsertToken(testestcaseDBPos.User.Email, testestcaseDBPos.token)
	if err != nil {
		t.Error(err)
	}
	err = conn.DropToken(testestcaseDBPos.token)
	if err != nil {
		t.Error(err)
	}
	truncateCustomTable("signed_users")
}
func TestDataBase_CheckEmail(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	err = conn.InsertToDB(testestcaseDBPos.User, "", "")
	err2 := conn.CheckMail(testestcaseDBPos.User.Email)
	if err2 {
		t.Error(err2)
	}
	truncateCustomTable("signed_users")
}
func TestDataBase_GetNames(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	err = conn.InsertToDB(testestcaseDBPos.User, "", "")
	_, _, err = conn.GetNames(testestcaseDBPos.User.Email)
	if err != nil {
		t.Error(err)
	}
	truncateCustomTable("signed_users")
}

func TestDataBase_CheckLoginPassword_failed(t *testing.T) {
	loginUser := models.LoginUser{
		LoginMail:     "123@123.1231",
		LoginPassword: "123123123",
	}
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
	err2 := conn.CheckLoginPassword(loginUser)
	if !err2 {
		t.Fatal("err is nil")
	}
	truncateCustomTable("signed_users")
}

func TestDataBase_InsertToDB_failed(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
	}
	testcaseDB_failed := testCase{
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
	t.Run("already exist", func(t *testing.T) {
		err = conn.InsertToDB(testcaseDB_failed.User, testcaseDB_failed.salt, testcaseDB_failed.hash)
		if err != nil {
			t.Error(err)
		}
		err = conn.InsertToDB(testcaseDB_failed.User, testcaseDB_failed.salt, testcaseDB_failed.hash)
		if err == nil {
			t.Fatal("err is nil")
		}
		truncateCustomTable("signed_users")
	})
	t.Run("empty value", func(t *testing.T) {
		testcaseDB_failed.User.FirstName = ""
		err = conn.InsertToDB(testcaseDB_failed.User, testcaseDB_failed.salt, testcaseDB_failed.hash)
		if err == nil {
			t.Fatal("err is nil")
		}
		truncateCustomTable("signed_users")

	})
}

func TestDataBase_InsertToken_failed(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	if err == nil {
		t.Fatal("err is nil")
	}
	truncateCustomTable("signed_users")
}
func TestDataBase_DropToken_failed(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	t.Run("no token", func(t *testing.T) {
		err = conn.DropToken(testestcaseDBPos.token)
		if err == nil {
			t.Fatal("err is nil")
		}
	})
	t.Run("different token", func(t *testing.T) {
		err = conn.InsertToDB(testestcaseDBPos.User, "", "")
		err = conn.InsertToken(testestcaseDBPos.User.Email, testestcaseDBPos.token)
		testestcaseDBPos.token = "qwerty"
		err = conn.DropToken(testestcaseDBPos.token)
		if err == nil {
			t.Fatal("err is nil")
		}
		truncateCustomTable("signed_users")
	})
}

func TestDataBase_CheckEmail_failed(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	t.Run("without email", func(t *testing.T) {
		err := conn.CheckMail(testestcaseDBPos.User.Email)
		if !err {
			t.Fatal("err is nil")
		}
		truncateCustomTable("signed_users")
	})
	t.Run("different email", func(t *testing.T) {
		err = conn.InsertToDB(testestcaseDBPos.User, "", "")
		testestcaseDBPos.User.Email = "123123123"
		err := conn.CheckMail(testestcaseDBPos.User.Email)
		if !err {
			t.Fatal("err is nil")
		}
		truncateCustomTable("signed_users")
	})
}

func TestDataBase_GetNames_failed(t *testing.T) {
	u := models.User{
		FirstName: "123",
		LastName:  "123",
		Email:     "123@123.1231",
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
	if err == nil {
		t.Fatal("err is nil")
	}
	truncateCustomTable("signed_users")
}
