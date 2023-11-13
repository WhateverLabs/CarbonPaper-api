package repository

import (
	"gorm.io/gorm"
)

type ApiRepository struct {
	DB *gorm.DB
}
