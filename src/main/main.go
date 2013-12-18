package main

import (
	"net/http"
	"server"
)

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/images/", http.FileServer(http.Dir("static")))

	server.RegisterHandlers()
	http.ListenAndServe(":80", nil)

}
