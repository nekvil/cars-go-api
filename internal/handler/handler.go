package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nekvil/cars-go-api/docs"
	"github.com/nekvil/cars-go-api/internal/service"
	"github.com/nekvil/cars-go-api/internal/utils"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SetupRouter() *gin.Engine {
	utils.Logger.Debug("Setting up router...")

	router := gin.Default()
	router.Use(LoggerMiddleware())

	v1 := router.Group("/v1")
	{
		v1.GET("/cars", h.GetAllCars)
		v1.DELETE("/cars/:id", h.DeleteCar)
		v1.PUT("/cars/:id", h.UpdateCar)
		v1.POST("/cars", h.AddCars)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	utils.Logger.Debug("Router setup completed")
	return router
}
