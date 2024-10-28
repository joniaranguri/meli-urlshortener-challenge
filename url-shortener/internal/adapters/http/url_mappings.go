package http

import (
	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/adapters/http/controllers"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/registry"
)

func defineRoutes(router gin.IRouter, appContainer registry.AppContainer) {
	// Shorten and retrieving endpoints
	urlController := controllers.UrlController{AppContainer: appContainer}
	urlController.RegisterRoutes(router)

	// Update and enable/disable endpoints
	urlManageController := controllers.UrlManageController{AppContainer: appContainer}
	urlManageController.RegisterRoutes(router)

}
