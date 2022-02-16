package main

import (
	"fmt"
	"log"
	"net/http"
	"registration-web-service2/pkg/service"
)

func loadMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "../assets/index.html")
}

func main() {
	http.HandleFunc("/", loadMainPage)
	http.HandleFunc("/signUp", service.SignUp)
	http.HandleFunc("/login", service.Login)
	http.HandleFunc("/sayname", service.SayName)
	http.HandleFunc("/logout", service.Logout)
	http.Handle("/login.html", http.FileServer(http.Dir("../assets")))
	http.Handle("/submit.html", http.FileServer(http.Dir("../assets")))
	http.Handle("/index.js", http.FileServer(http.Dir("../assets")))
	http.Handle("/eNWDJx.jpg", http.FileServer(http.Dir("../assets")))
	http.Handle("/style.css", http.FileServer(http.Dir("../assets")))
	fmt.Printf("Starting server for testing HTTP POST in 8081...\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
