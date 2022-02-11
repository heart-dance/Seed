package app

import (
	"github.com/heart-dance/seed/src/app/server"
	"go.uber.org/zap"
)

type Application struct {
	version string
	srv     server.HttpServer
	logger  *zap.Logger
}

func NewApplication(version, host, profile, web string) (*Application, error) {
	if err := CheckProfile(profile); err != nil {
		return nil, err
	}

	NewDB(profile)

	logger := NewLogger("dev", "./test/app/data/log")
	logger.Info("Starting application")
	return &Application{
		version: version,
		srv:     server.NewHttpServer(),
		logger:  logger,
	}, nil
}

func (a *Application) Run() error {
	a.srv.Run()
	return nil
}
