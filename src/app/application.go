package app

import (
	"github.com/ferralucho/store_users-api/src/datasources/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("starting the application")
	router.Run(":8082")
}
