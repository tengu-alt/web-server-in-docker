package handlers

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	us := models.LoginUser{}
	err = json.Unmarshal(data, &us)
	if err != nil {
		return
	}
	resp, err := service.Login(us)
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

func Logout(w http.ResponseWriter, r *http.Request) {
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
	err = service.Logout(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func SayNameHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	resp := TokenResponse{}
	Message, time := service.SayName(token)
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
