package app

import (
	"path/filepath"

	"github.com/heart-dance/seed/src/app/db"
	"github.com/heart-dance/seed/src/app/server"
	"go.uber.org/zap"
)

type Application struct {
	version string
	srv     server.HttpServer
	logger  *zap.Logger
	db      db.DB
}

func NewApplication(version, host, profile, web, runMode string) (*Application, error) {
	if err := CheckProfile(profile); err != nil {
		return nil, err
	}
	logPath := filepath.Join(profile, "log")

	logger := NewLogger(runMode, logPath)
	logger.Info("Starting application.")
	logger.Info("Application version: " + version + " run mode: " + runMode)
	db := db.NewDB(version, profile, logger)

	return &Application{
		version: version,
		srv:     server.NewHttpServer(logger, db),
		logger:  logger,
		db:      db,
	}, nil
}

func (a *Application) Run() error {
	a.srv.Run()
	return nil
}
