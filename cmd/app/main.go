package main

import (
	"link_service/internal/config"
	"link_service/internal/db"
	"link_service/internal/handler"
	"link_service/internal/link_service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config := config.GetConfig()

	storage := db.NewStorage()

	linkService := link_service.NewService(config, storage)

	handler := handler.NewHandler(router, linkService, config)
	handler.Register()

	router.Run(config.Addr)
}
