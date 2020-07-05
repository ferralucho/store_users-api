package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/store_users-api/datasources/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Log.Info("starting the application")
	router.Run(":8080")
}
