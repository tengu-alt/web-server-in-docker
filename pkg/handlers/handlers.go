package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"registration-web-service2/pkg/models"
	"registration-web-service2/pkg/service"
	"registration-web-service2/pkg/store"
	"registration-web-service2/pkg/validation"
)

type User = models.User
type ValidationErr = models.ValidationErr
type LoginUser = models.LoginUser
type TokenResponse = models.TokenResponse

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	u := *&models.User{}
	err = json.Unmarshal(data, &u)
	if err != nil {
		return
	}
	validationErrors, err := service.SignUp(u)
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	us := *&models.LoginUser{}
	err = json.Unmarshal(data, &us)
	if err != nil {
		return
	}
	resp := TokenResponse{}
	if validation.LoginValid(us) == true {
		token := service.GiveToken(us)
		resp.ResponseMessage = "success login"
		resp.Token = token
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	} else {
		resp.ResponseMessage = "invalid data"
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	var token string
	err = json.Unmarshal(data, &token)
	if err != nil {
		return
	}
	store.DropToken(token)
	w.Write(b)
}

func SayName(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	resp := TokenResponse{}
	Message, time := validation.SayNameFunc(token)
	if !time {
		resp.ResponseMessage = "login again"
		store.DropToken(token)
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
