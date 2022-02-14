package db

func (d *db) GetWebAuthUser() string {
	return d.webConfigData.WebAuthUser
}

func (d *db) GetWebAuthPwd() string {
	return d.webConfigData.WebAuthPwd
}
