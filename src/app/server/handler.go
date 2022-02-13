package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/heart-dance/seed/src/app/db"
	"go.uber.org/zap"
)

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Info(w http.ResponseWriter, r *http.Request)
	GetConfig(w http.ResponseWriter, r *http.Request)
	UpdateConfig(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger *zap.Logger
	db     db.DB
}

func NewHandler(logger *zap.Logger, db db.DB) Handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("add"))
}

func (h *handler) Info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("info"))
}

func (h *handler) GetConfig(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()
	h.logger.Debug(fmt.Sprintf("get key %v", query.Get("key")))
	w.Header().Set("Content-Type", "application/json")
	if query.Get("key") == "common_config" {
		var data = h.db.GetCommonConfigData()
		json.NewEncoder(w).Encode(data)
	} else if query.Get("key") == "web_config" {
		var data = h.db.GetWebConfigData()
		json.NewEncoder(w).Encode(data)
	}
}

func (h *handler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updateconfig"))
}
