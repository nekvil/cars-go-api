package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nekvil/cars-go-api/internal/handler"
	"github.com/nekvil/cars-go-api/internal/repository"
	"github.com/nekvil/cars-go-api/internal/service"
	"github.com/nekvil/cars-go-api/internal/utils"
	"github.com/nekvil/cars-go-api/server"
)

// @title Cars Go API
// @version 0.0.1
// @description This API provides endpoints for managing cars.

// @host localhost:8080
// @BasePath /v1
func main() {
	utils.InitLogger()

	db := repository.SetupDatabase()
	repos := repository.NewRepository(db, utils.GetEnv("EXTERNAL_API_URL"))
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	utils.Logger.Info("Starting server...")
	server := new(server.Server)
	port := ":" + utils.GetEnv("PORT")
	go func() {

		if err := server.Run(port, handlers.SetupRouter()); err != nil {
			utils.Logger.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	utils.Logger.Info("The server is shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		utils.Logger.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
