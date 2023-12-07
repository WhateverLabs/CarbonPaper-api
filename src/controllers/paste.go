package controllers

import (
	"carbon-paper/src/repository"
	"carbon-paper/src/types"
	"log"

	"github.com/gin-gonic/gin"
)

type PasteController struct {
	repo *repository.PasteRepository
}

func NewPasteController(repo *repository.PasteRepository) *PasteController {
	return &PasteController{
		repo: repo,
	}
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

	newID, err := ctrl.repo.CreatePaste(&body, body.ExpiresInSeconds)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, types.GenericResponseBody{
			Success: false,
			Message: "Could not save paste",
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

	paste, err := ctrl.repo.GetPasteByID(pasteID)
	if err != nil {
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

	paste, err := ctrl.repo.GetPasteByID(pasteID)
	if err != nil {
		c.AbortWithStatusJSON(404, types.GenericResponseBody{
			Success: false,
			Message: "Paste not found",
		})
		return
	}

	if paste.PasteRequestBody.OneView {
		ctrl.repo.DeletePasteByID(pasteID)
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
