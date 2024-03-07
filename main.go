package main

import (
	"example_site/config"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	Domain string `env:"DOMAIN"`
}

func (s *Server) redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+s.Domain+":443"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	cfg := config.New()

	server := &Server{
		Domain: cfg.Domain,
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Redirect HTTP to HTTPS
	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(server.redirectTLS)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, ""); err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(400)
		}
		w.WriteHeader(200)
	})

	log.Fatal(http.Serve(autocert.NewListener(server.Domain), r))
}
