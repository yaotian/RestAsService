package server

import (
	"data"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)

type Config struct {
}

const PathPrefix = "/"

// badRequest is handled by setting the status code in the reply to StatusBadRequest.
type badRequest struct{ error }

// notFound is handled by setting the status code in the reply to StatusNotFound.
type notFound struct{ error }

// errorHandler wraps a function returning an error by handling the error and returning a http.Handler.
// If the error is of the one of the types defined above, it is handled as described for every type.
// If the error is of another type, it is considered as an internal error and its message is logged.
func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "not found", http.StatusNotFound)
		default:
			logger.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix, errorHandler(ListAll)).Methods("GET")
	r.HandleFunc(PathPrefix, errorHandler(New)).Methods("POST")
	r.HandleFunc(PathPrefix+"{id}", errorHandler(GetOne)).Methods("GET")
	r.HandleFunc(PathPrefix+"{id}", errorHandler(UpdateOne)).Methods("PUT")
	r.HandleFunc("/help/", Help)
	r.HandleFunc("/admin/", Admin)
	http.Handle(PathPrefix, r)

}

var dataMgr = data.NewDataManager()

func ListAll(w http.ResponseWriter, r *http.Request) error {
	res := struct{ Items []*data.Item }{dataMgr.All()}
	return json.NewEncoder(w).Encode(res)
}

func New(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func GetOne(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func UpdateOne(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func Help(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/help.html")
	if err != nil {
		return
	}
	t.Execute(w, nil)
}
func Admin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		return
	}
	t.Execute(w, nil)
}
