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

func (ctrl *PasteController) GetPasteMetadata(c *gin.Context) {
	pasteID := c.Param("pasteID")

	paste := models.Paste{}
	if err := ctrl.DB.Where("id = ?", pasteID).First(&paste).Error; err != nil {
		c.AbortWithStatusJSON(404, types.GenericResponseBody{
			Success: false,
			Message: "Paste not found",
		})
		return
	}

	c.JSON(200, types.GenericResponseBody{
		Success: true,
		Message: "Paste metadata was found",
		Data: gin.H{
			"passwordHashSalt": paste.PasteRequestBody.PasswordHashSalt,
		},
	})
}

func (ctrl *PasteController) GetPaste(c *gin.Context) {
	pasteID := c.Param("pasteID")

	paste := models.Paste{}
	if err := ctrl.DB.Where("id = ?", pasteID).First(&paste).Error; err != nil {
		c.AbortWithStatusJSON(404, types.GenericResponseBody{
			Success: false,
			Message: "Paste not found",
		})
		return
	}

	if paste.PasteRequestBody.OneView {
		ctrl.DB.Delete(&paste)
	}

	if paste.PasteRequestBody.KekHash != "" {
		providedHash := c.Query("passwordHash")

		if providedHash != paste.PasteRequestBody.KekHash {
			c.AbortWithStatusJSON(401, types.GenericResponseBody{
				Success: false,
				Message: "Incorrect password",
			})
			return
		}
	}

	c.JSON(200, types.GenericResponseBody{
		Success: true,
		Message: "Paste was found",
		Data:    paste.PasteRequestBody,
	})
}
