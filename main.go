package main

import (
	"example_site/config"
	"example_site/logger"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/acme/autocert"
	"html/template"
	"net/http"
)

type Server struct {
	config *config.Config
	tmpl   *template.Template
	log    *zerolog.Logger
}

func (s *Server) redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+s.config.Domain+":443"+r.RequestURI, http.StatusMovedPermanently)
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	if err := s.tmpl.Execute(w, ""); err != nil {
		s.log.Error().Err(err).Send()
		w.Write([]byte(err.Error()))
		w.WriteHeader(400)
	}
	w.WriteHeader(200)
}

func main() {
	cfg := config.New()
	logg := logger.New(cfg.LogLevel)
	tmpl, err := template.ParseFiles("template/startbootstrap-freelancer-gh-pages/index.html")
	if err != nil {
		logg.Fatal().Err(err).Send()
	}

	server := &Server{
		config: cfg,
		tmpl:   tmpl,
		log:    logg,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", server.Home).Methods(http.MethodGet)
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./template/startbootstrap-freelancer-gh-pages/static/")))
	r.PathPrefix("/static/").Handler(s)
	if server.config.CreateTLSConfig {
		// Redirect HTTP to HTTPS
		go func() {
			if err := http.ListenAndServe(":80", http.HandlerFunc(server.redirectTLS)); err != nil {
				server.log.Fatal().Err(err).Send()
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
		if err := s.ListenAndServeTLS("", ""); err != nil {
			server.log.Fatal().Err(err).Send()
		}
	}

	if err := http.ListenAndServe(":80", r); err != nil {
		server.log.Fatal().Err(err).Send()
	}
}
