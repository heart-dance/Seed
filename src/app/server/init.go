package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer interface {
	Run() error
}

type httpServe struct {
	srv *http.Server
}

func NewHttpServer() HttpServer {
	r := mux.NewRouter()
	h := NewHandler()
	r.Use(AuthMdiddleware("123", "123").Middleware)
	s := r.PathPrefix("/api/v1").Subrouter()
	// s.HandleFunc("/login", h.Login).Methods("POST")
	s.HandleFunc("/task", h.Add).Methods("POST")
	s.HandleFunc("/task", h.Info).Methods("GET")
	s.HandleFunc("/config", h.GetConfig).Methods("GET")
	s.HandleFunc("/config", h.UpdateConfig).Methods("PUT")

	return &httpServe{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
	}
}

func (s *httpServe) Run() error {
	s.srv.ListenAndServe()
	return nil
}
