package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, ""); err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(400)
		}
		w.WriteHeader(200)
	})
	log.Fatal(http.ListenAndServe(":80", r))
}
