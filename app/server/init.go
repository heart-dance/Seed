package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heart-dance/seed/app/db"
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
	m := NewMiddleware(logger, db)
	r.Use(m.LoggingMiddleware)
	r.Use(m.AuthMiddleware)
	s := r.PathPrefix("/api/v1").Subrouter()
	// s.HandleFunc("/login", h.Login).Methods("POST")
	s.HandleFunc("/task", h.Add).Methods("POST")
	s.HandleFunc("/task", h.Info).Methods("GET")
	s.HandleFunc("/config", h.GetConfig).Methods("GET")
	s.HandleFunc("/config", h.UpdateConfig).Methods("PUT")

	if db.GetWebConfigData().WebUIPath != "" {
		logger.Debug("Http server enable webui, path:" + db.GetWebConfigData().WebUIPath)
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(db.GetWebConfigData().WebUIPath))))
	}

	return &httpServe{
		logger: logger,
		srv: &http.Server{
			Addr:    db.GetWebConfigData().WebHost,
			Handler: r,
		},
	}
}

func (s *httpServe) Run() error {
	s.logger.Debug("Starting http server, listening on port: " + s.srv.Addr)
	s.srv.ListenAndServe()
	return nil
}
