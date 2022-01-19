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

func Login(w http.ResponseWriter, r *http.Request) {
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
