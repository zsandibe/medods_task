package app

import (
	"task/internal/config"
	"task/internal/delivery"
	"task/internal/service"
	"task/migration"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	router := gin.Default()

	configs := config.Configuration()

	mongoDb := migration.NewDatabase(configs)

	storage := migration.NewCollections(mongoDb, configs.NameCollection)

	authService := service.NewService(storage)

	handler := delivery.NewHandlers(authService)

	handler.Routes(router)

	return router
}
