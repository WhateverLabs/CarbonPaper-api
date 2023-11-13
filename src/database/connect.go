package database

import (
	"carbon-paper/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase(name string) *gorm.DB {
	gormConfig := &gorm.Config{}

	db, err := gorm.Open(sqlite.Open(name), gormConfig)
	if err != nil {
		panic("Could not connect to SQLITE database. Error: " + err.Error())
	}

	migrations(db)

	return db
}

func migrations(db *gorm.DB) {
	db.AutoMigrate(&models.Paste{})
}
