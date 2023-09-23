package api

import "gvb/api/setting_api"

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
