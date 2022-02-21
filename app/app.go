package app

import (
	"path/filepath"

	"github.com/heart-dance/seed/app/db"
	"github.com/heart-dance/seed/app/server"
	"go.uber.org/zap"
)

type Application struct {
	logger *zap.Logger
	db     db.DB
	srv    server.HttpServer
}

func NewApplication(version, host, profile, web, runMode string) (*Application, error) {
	if err := CheckProfile(profile); err != nil {
		return nil, err
	}
	logPath := filepath.Join(profile, "log")

	logger := NewLogger(runMode, logPath)
	logger.Debug("Starting application.")
	logger.Debug("Application version: " + version + " run mode: " + runMode)
	db := db.NewDB(version, profile, host, web, logger)

	return &Application{
		logger: logger,
		db:     db,
		srv:    server.NewHttpServer(logger, db),
	}, nil
}

func (a *Application) Run() error {
	a.srv.Run()
	return nil
}
