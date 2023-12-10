package controllers

import (
	"carbon-paper/src/repository"

	"github.com/gin-gonic/gin"
)

type MiddlewareController struct {
	RatelimitRepository *repository.RatelimitRepository
}

func NewMiddlewareController(repo *repository.RatelimitRepository) *MiddlewareController {
	return &MiddlewareController{
		RatelimitRepository: repo,
	}
}

func (ctrl *MiddlewareController) RatelimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ctrl.RatelimitRepository.Get(c.ClientIP()) > 10 {
			c.AbortWithStatusJSON(429, gin.H{
				"message": "Too many requests",
			})
			return
		}

		ctrl.RatelimitRepository.Increment(c.ClientIP())

		c.Next()
	}
}
