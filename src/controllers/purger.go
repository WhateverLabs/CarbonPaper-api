package controllers

import (
	"carbon-paper/src/models"
	"carbon-paper/src/repository"
	"time"
)

type PurgerController struct {
	repository.ApiRepository
}

func (ctrl *PurgerController) Purge() {
	ctrl.DB.Where("expires_at < ?", time.Now()).Delete(&models.Paste{})
}
