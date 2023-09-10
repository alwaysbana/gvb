package global

import (
	"gorm.io/gorm"
	"gvb/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
