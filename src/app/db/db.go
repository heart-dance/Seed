package db

import (
	"path/filepath"

	"github.com/timshannon/bolthold"
	"go.uber.org/zap"
)

type DB interface {
	GetCommonConfig() (commonConfigData, error)
	GetWebConfig() (webConfigData, error)
	GetCommonConfigData() commonConfigData
	GetWebConfigData() webConfigData
}

type db struct {
	logger           *zap.Logger
	ConfigDB         *bolthold.Store
	commonConfigData commonConfigData
	webConfigData    webConfigData
}

func NewDB(version, profile, host, webPath string, logger *zap.Logger) DB {
	configPath := filepath.Join(profile, "/config.db")
	confDB, _ := bolthold.Open(configPath, 0666, nil)

	newDB := &db{
		logger:   logger,
		ConfigDB: confDB,
	}
	_ = newDB.initData(version, profile, host, webPath)
	return newDB
}

type commonConfigData struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	ProfilePath string `json:"profile_path"`
}

type webConfigData struct {
	WebHost   string `json:"web_host"`
	WebUIPath string `json:"web_ui_path"`
}

func (d *db) initData(version, profile, host, webPath string) error {
	if err := d.initCommonConfigData(version, profile); err != nil {
		return err
	}
	if err := d.initWebConfigData(host, webPath); err != nil {
		return err
	}

	return nil
}

func (d *db) initCommonConfigData(version, profile string) error {
	var initData = commonConfigData{
		Name:        "seed",
		Version:     version,
		ProfilePath: profile,
	}
	var data commonConfigData
	err := d.ConfigDB.Get("common_config", &data)
	if err != nil {
		if bolthold.ErrNotFound == err {
			d.logger.Debug("init common config data")
			d.ConfigDB.Insert("common_config", initData)
			d.commonConfigData = initData
			return nil
		} else {
			d.logger.Error("init common config data error", zap.Error(err))
			return err
		}
	} else {
		d.commonConfigData = data
		return nil
	}
}

func (d *db) initWebConfigData(host, path string) error {
	var initData = webConfigData{
		WebHost:   host,
		WebUIPath: path,
	}
	var data webConfigData
	err := d.ConfigDB.Get("web_config", &data)
	if err != nil {
		if bolthold.ErrNotFound == err {
			d.logger.Debug("init web config data")
			d.ConfigDB.Insert("web_config", initData)
			d.webConfigData = initData
			return nil
		} else {
			d.logger.Error("init common config data error", zap.Error(err))
			return err
		}
	} else {
		d.webConfigData = data
		return nil
	}
}

func (d *db) GetCommonConfigData() commonConfigData {
	return d.commonConfigData
}

func (d *db) GetWebConfigData() webConfigData {
	return d.webConfigData
}
