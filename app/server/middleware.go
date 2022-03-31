package server

import (
	"net/http"

	"github.com/heart-dance/seed/app/db"
	"go.uber.org/zap"
)

type Middleware interface {
	AuthMiddleware(next http.Handler) http.Handler
	LoggingMiddleware(next http.Handler) http.Handler
}
type middleware struct {
	logger *zap.Logger
	db     db.DB
}

func NewMiddleware(logger *zap.Logger, configDB db.DB) Middleware {
	return &middleware{
		logger: logger,
		db:     configDB,
	}
}

func (m *middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUser, reqPasswd, hasAuth := r.BasicAuth()
		if (m.db.GetWebAuthUser() == "" && m.db.GetWebAuthPwd() == "") ||
			(hasAuth && reqUser == m.db.GetWebAuthUser() && reqPasswd == m.db.GetWebAuthPwd()) {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}

func (m *middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Debug("Request: " + r.Method + " " + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
