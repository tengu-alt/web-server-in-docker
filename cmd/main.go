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
	connect := handlers.NewConnect(connString)
	http.HandleFunc("/", loadMainPage)
	http.HandleFunc("/signUp", handlers.NewSignUpHandler(connect))
	http.HandleFunc("/login", handlers.NewLoginHandler(connect))
	http.HandleFunc("/sayname", handlers.NewSayNameHandler(connect))
	http.HandleFunc("/logout", handlers.NewLogoutHandler(connect))
	fmt.Printf("Starting server for testing HTTP POST in 8081...\n")
	if err := http.ListenAndServe("0.0.0.0:8081", nil); err != nil {
		log.Fatal(err)
	}
}
