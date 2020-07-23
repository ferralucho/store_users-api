package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ferralucho/store_users-api/datasources/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("starting the application")
	router.Run(":8082")
}
