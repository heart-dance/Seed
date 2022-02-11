package app

import (
	"github.com/heart-dance/seed/src/app/server"
	"go.uber.org/zap"
)

type Application struct {
	srv    server.HttpServer
	logger *zap.Logger
}

func NewApplication(host, profile, web string) (*Application, error) {
	logger := NewLogger("dev", "./test/app/data/log")
	logger.Info("Starting application")
	return &Application{
		srv:    server.NewHttpServer(),
		logger: logger,
	}, nil
}

func (a *Application) Run() error {
	a.srv.Run()
	return nil
}
