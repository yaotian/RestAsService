package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Config struct {
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/images/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", configHandler)
	http.ListenAndServe(":8888", nil)
}
