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
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/add", h.Login).Methods("POST")
	r.HandleFunc("/info", h.Login).Methods("GET")

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
