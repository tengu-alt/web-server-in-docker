package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"registration-web-service2/pkg/store"
	"registration-web-service2/pkg/users"
	"registration-web-service2/pkg/validation"
)

type User = users.User
type ValidationErr = users.ValidationErr
type LoginUser = users.LoginUser

func SignUp(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	u := *&users.User{}
	err = json.Unmarshal(data, &u)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	validationErrors := validation.Validate(u)
	fmt.Printf("%s", validationErrors)
	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)

		b, err := json.Marshal(&validationErrors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
		return
	}
	fmt.Print("user", u)
	store.InsertToDB(u)
	w.Write([]byte("[{}]"))
}
func Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()
	data := []byte(b)
	us := *&users.LoginUser{}
	err = json.Unmarshal(data, &us)
	if err != nil {
		return
	}
	resp := ValidationErr{}
	if validation.LoginValid(us) == true {
		resp.ErrMassage = "success login"
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	} else {
		resp.ErrMassage = "invalid data"
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}

}
