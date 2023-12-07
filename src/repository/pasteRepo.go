package repository

import (
	"carbon-paper/src/models"
	"carbon-paper/src/types"
	"carbon-paper/src/utils"
	"time"

	"gorm.io/gorm"
)

type PasteRepository struct {
	DB *gorm.DB
}

func NewPasteRepository(db *gorm.DB) *PasteRepository {
	return &PasteRepository{
		DB: db,
	}
}

func (repo *PasteRepository) CreatePaste(body *types.PasteRequestBody, expiresInSeconds int) (newID string, err error) {
	newID, err = utils.GenerateRandomID(10)
	if err != nil {
		return
	}

	paste := models.Paste{
		ID:               newID,
		PasteRequestBody: *body,
		CreatedAt:        time.Now(),
		ExpiresAt:        time.Now().Add(time.Duration(expiresInSeconds) * time.Second),
	}
	err = repo.DB.Create(&paste).Error
	return
}

func (repo *PasteRepository) GetPasteByID(pasteID string) (paste *models.Paste, err error) {
	paste = &models.Paste{}
	err = repo.DB.Where("id = ?", pasteID).First(&paste).Error
	return
}

func (repo *PasteRepository) DeletePasteByID(pasteID string) (err error) {
	err = repo.DB.Where("id = ?", pasteID).Delete(&models.Paste{}).Error
	return
}
