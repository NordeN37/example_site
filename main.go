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
	config *config.Config
	tmpl   *template.Template
}

func (s *Server) redirectTLS(w http.ResponseWriter, r *http.Request) {
	log.Println("Domain: ", s.config.Domain)
	http.Redirect(w, r, "https://"+s.config.Domain+":443"+r.RequestURI, http.StatusMovedPermanently)
}

func (s *Server) CallbackAnyMoney(w http.ResponseWriter, r *http.Request) {
	log.Println("CallbackAnyMoney ")
	w.WriteHeader(200)
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	if err := s.tmpl.Execute(w, ""); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
	}
	w.WriteHeader(200)
}

func main() {
	cfg := config.New()

	tmpl, err := template.ParseFiles("template/startbootstrap-freelancer-gh-pages/index.html")
	if err != nil {
		log.Fatal(err)
	}

	server := &Server{
		config: cfg,
		tmpl:   tmpl,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", server.Home).Methods(http.MethodGet)
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./template/startbootstrap-freelancer-gh-pages/static/")))
	r.PathPrefix("/static/").Handler(s)

	if server.config.CreateTLSConfig {
		// Redirect HTTP to HTTPS
		go func() {
			if err := http.ListenAndServe(":80", http.HandlerFunc(server.redirectTLS)); err != nil {
				log.Fatalf("ListenAndServe error: %v", err)
			}
		}()
		m := &autocert.Manager{
			Cache:      autocert.DirCache("secret-dir"),
			Prompt:     autocert.AcceptTOS,
			Email:      server.config.Email,
			HostPolicy: autocert.HostWhitelist(server.config.Domain),
		}
		s := &http.Server{
			Addr:      ":https",
			Handler:   r,
			TLSConfig: m.TLSConfig(),
		}
		log.Fatal(s.ListenAndServeTLS("", ""))
	}
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
