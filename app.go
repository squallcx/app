package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/ip", ipHandler)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/s", shortHandler)
	http.Handle("/", r)
}
