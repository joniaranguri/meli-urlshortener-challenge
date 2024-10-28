package http

import (
	"fmt"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/registry"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/config"
	cors "github.com/rs/cors/wrapper/gin"
)

func Start() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "81"
	}

	if err := run(port); err != nil {
		fmt.Printf("error running server %s", err)
	}
}

func run(port string) error {
	if os.Getenv("SCOPE") != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	s, _ := new()

	return s.Run(fmt.Sprintf(":%s", port))
}

func new() (*gin.Engine, *registry.AppContainer) {

	var router *gin.Engine
	if os.Getenv("SCOPE") != "local" {
		router = gin.Default()
	} else {
		router = gin.New()

	}

	router.NoRoute(cors.AllowAll())

	r := registry.NewRegistry(configs.Conf)

	appContainer := r.NewAppContainer()

	defineRoutes(router, appContainer)

	addHealthHandler(router)

	return router, &appContainer
}

func addHealthHandler(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
}
