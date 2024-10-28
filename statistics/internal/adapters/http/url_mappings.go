package http

import (
	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/adapters/http/controllers"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/registry"
)

func defineRoutes(router gin.IRouter, appContainer registry.AppContainer) {
	// Save(mock) and retrieve statistics endpoints
	StatisticsController := controllers.StatisticsController{AppContainer: appContainer}
	StatisticsController.RegisterRoutes(router)

}
