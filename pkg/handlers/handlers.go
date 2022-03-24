package handlers

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"web-server-in-docker/pkg/models"
	"web-server-in-docker/pkg/service"
	"web-server-in-docker/pkg/store"
)

type User = models.User
type ValidationErr = models.ValidationErr
type LoginUser = models.LoginUser
type TokenResponse = models.TokenResponse
type DBconn struct {
	Data *store.DataBase
}

func NewConnStruct(db *sqlx.DB) *DBconn {
	return &DBconn{
		Data: &store.DataBase{
			DBmodel: db,
		},
	}

}

func NewSignUpHandler(db *DBconn) http.HandlerFunc {
	return db.SignUpHandler

}

func NewLoginHandler(db *DBconn) http.HandlerFunc {
	return db.LoginHandler

}

func NewLogoutHandler(db *DBconn) http.HandlerFunc {
	return db.LogoutHandler

}

func NewSayNameHandler(db *DBconn) http.HandlerFunc {
	return db.SayNameHandler

}

func (connectStruct *DBconn) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	u := User{}
	err = json.Unmarshal(data, &u)
	if err != nil {
		return
	}
	validationErrors, err := service.SignUp(u, connectStruct.Data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		b, err := json.Marshal(&validationErrors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
		return
	}
	w.Write([]byte("[{}]"))
}

func (connectStruct *DBconn) LoginHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	us := LoginUser{}
	err = json.Unmarshal(data, &us)
	if err != nil {
		return
	}
	resp, err := service.Login(us, connectStruct.Data)
	if err != nil {
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)

	} else {
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}

}

func (connectStruct *DBconn) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	var token string
	err = json.Unmarshal(data, &token)
	if err != nil {
		return
	}
	err = service.Logout(token, connectStruct.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (connectStruct *DBconn) SayNameHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	resp := TokenResponse{}
	Message, time := service.SayName(token)
	if !time {
		resp.ResponseMessage = "login again"
		connectStruct.Data.DropToken(token)
	} else {
		resp.ResponseMessage = Message
	}
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
