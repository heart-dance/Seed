package db

import "github.com/heart-dance/seed/app/model"

func (d *db) GetCommonConfig() (commonConfigData, error) {
	var data commonConfigData
	d.ConfigDB.Get("common_config", &data)
	return data, nil
}

func (d *db) UpdateCommonConfig() (model.Config, error) {
	var data model.Config
	d.ConfigDB.Get("common_config", &data)
	return data, nil
}

func (d *db) GetWebConfig() (webConfigData, error) {
	var data webConfigData
	d.ConfigDB.Get("web_config", &data)
	return data, nil
}
