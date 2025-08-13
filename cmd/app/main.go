package main

import (
	"link_service/internal/config"
	"link_service/internal/db"
	"link_service/internal/handler"
	"link_service/internal/link_service"
	"link_service/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.GetLogger()

	router := gin.Default()

	config := config.GetConfig(logger)

	storage := db.NewStorage(logger)

	linkService := link_service.NewService(config, storage, logger)

	handler := handler.NewHandler(router, linkService, config)
	handler.Register()

	router.Run(config.Addr)
}
