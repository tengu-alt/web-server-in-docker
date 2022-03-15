package main

import (
	"fmt"
	"log"
	"net/http"
	"web-server-in-docker/pkg/handlers"
	"web-server-in-docker/pkg/store"
)

func loadMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "../assets/index.html")
		return
	}

	http.ServeFile(w, r, "../assets/"+r.URL.Path)
}

func main() {
	connString := store.GetConfig()
	database, err := store.NewConnect(connString)
	if err != nil {
		log.Fatal(err)
	}
	connection := handlers.NewConnStruct(database)
	http.HandleFunc("/", loadMainPage)
	http.HandleFunc("/signUp", handlers.NewSignUpHandler(connection))
	http.HandleFunc("/login", handlers.NewLoginHandler(connection))
	http.HandleFunc("/sayname", handlers.NewSayNameHandler(connection))
	http.HandleFunc("/logout", handlers.NewLogoutHandler(connection))
	fmt.Printf("Starting server for testing HTTP POST in 8081...\n")
	if err := http.ListenAndServe("0.0.0.0:8081", nil); err != nil {
		log.Fatal(err)
	}
}
