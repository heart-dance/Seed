package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heart-dance/seed/src/app/db"
	"go.uber.org/zap"
)

type HttpServer interface {
	Run() error
}

type httpServe struct {
	srv    *http.Server
	logger *zap.Logger
}

func NewHttpServer(logger *zap.Logger, db db.DB) HttpServer {
	r := mux.NewRouter()
	h := NewHandler(logger, db)
	r.Use(AuthMdiddleware("123", "123").Middleware)
	s := r.PathPrefix("/api/v1").Subrouter()
	// s.HandleFunc("/login", h.Login).Methods("POST")
	s.HandleFunc("/task", h.Add).Methods("POST")
	s.HandleFunc("/task", h.Info).Methods("GET")
	s.HandleFunc("/config", h.GetConfig).Methods("GET")
	s.HandleFunc("/config", h.UpdateConfig).Methods("PUT")

	return &httpServe{
		logger: logger,
		srv: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
	}
}

func (s *httpServe) Run() error {
	s.logger.Info("Starting http server")
	s.srv.ListenAndServe()
	return nil
}
