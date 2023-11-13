package repository

import (
	"carbon-paper/src/config"

	"gorm.io/gorm"
)

type ApiRepository struct {
	DB  *gorm.DB
	Cfg *config.Config
}
