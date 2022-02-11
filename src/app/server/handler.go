package server

import "net/http"

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}
