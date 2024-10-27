package http

import (
	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/adapters/http/controllers"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/registry"
)

func defineRoutes(router gin.IRouter, appContainer registry.AppContainer) {
	endpointController := controllers.UrlController{AppContainer: appContainer}
	endpointController.RegisterRoutes(router)
}
