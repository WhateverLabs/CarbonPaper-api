package main

import (
	"carbon-paper/src/config"
	"carbon-paper/src/controllers"
	"carbon-paper/src/database"
	"carbon-paper/src/repository"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Parse()

	db := database.ConnectDatabase(cfg.DatabaseName)

	// set logging location
	f, err := os.OpenFile(fmt.Sprintf("%s/carbon-paper-%d.log", cfg.LogLocation, time.Now().Unix()), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	r := gin.Default()
	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:  cfg.CorsAllowOrigins,
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	apiRepo := repository.ApiRepository{
		DB:  db,
		Cfg: cfg,
	}

	pasteController := controllers.PasteController{
		ApiRepository: apiRepo,
	}

	r.POST("/new", pasteController.CreatePaste)
	r.GET("/metadata/:pasteID", pasteController.GetPasteMetadata)
	r.GET("/data/:pasteID", pasteController.GetPaste)

	r.Run(fmt.Sprintf("%s:%s", cfg.ListenHost, cfg.ListenPort))
}
