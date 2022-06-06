package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"test_project_sell/config"
	"test_project_sell/internal/handler"
	"test_project_sell/internal/repository"
	"test_project_sell/internal/server"
	"test_project_sell/internal/service"
)

func main() {
	cfg := config.NewConfig()
	server := server.Server{}
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	gin.SetMode(gin.ReleaseMode)

	log.Printf("Server starting on post %s...", cfg.Port)
	if err := server.Start(":"+cfg.Port, handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
