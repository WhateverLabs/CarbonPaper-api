package controllers

import (
	"carbon-paper/src/models"
	"carbon-paper/src/repository"
	"carbon-paper/src/types"
	"carbon-paper/src/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type PasteController struct {
	repository.ApiRepository
}

func (ctrl *PasteController) CreatePaste(c *gin.Context) {
	body := types.PasteRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, types.GenericResponseBody{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if body.PasswordHash != "" && body.PasswordHashSalt == "" {
		c.AbortWithStatusJSON(400, types.GenericResponseBody{
			Success: false,
			Message: "Password hash provided but no salt",
		})
	}

	if body.PasswordHash == "" && body.PasswordHashSalt != "" {
		c.AbortWithStatusJSON(400, types.GenericResponseBody{
			Success: false,
			Message: "Password salt provided but no hash",
		})
	}

	newID, err := utils.GenerateRandomID(10)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, types.GenericResponseBody{
			Success: false,
			Message: "Unable to generate ID",
		})
		return
	}

	paste := models.Paste{
		ID:               newID,
		PasteRequestBody: body,
		CreatedAt:        time.Now(),
		ExpiresAt:        time.Now().Add(time.Duration(body.ExpiresInSeconds) * time.Second),
	}
	if err := ctrl.DB.Create(&paste).Error; err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, types.GenericResponseBody{
			Success: false,
			Message: "Unable to save paste",
		})
		return
	}

	c.JSON(200, types.GenericResponseBody{
		Success: true,
		Message: "Paste was saved",
		Data:    newID,
	})
}
