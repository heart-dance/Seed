package server

import "net/http"

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Info(w http.ResponseWriter, r *http.Request)
	GetConfig(w http.ResponseWriter, r *http.Request)
	UpdateConfig(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
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
	w.Write([]byte("getconfig"))
}

func (h *handler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updateconfig"))
}
