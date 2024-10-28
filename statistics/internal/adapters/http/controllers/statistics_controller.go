package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/registry"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/utils"
	"net/http"
	"os"
)

type StatisticsController struct {
	AppContainer registry.AppContainer
}

func (controller *StatisticsController) RegisterRoutes(router gin.IRouter) {
	if os.Getenv("SCOPE") == "DEMO" {
		router.POST("/:shortUrlId/trackClick", controller.PostTrackClick)
	}
	router.GET("/:shortUrlId", controller.GetClickStatistics)
}

func (controller *StatisticsController) PostTrackClick(c *gin.Context) {
	shortUrlId := c.Param("shortUrlId")

	err := controller.AppContainer.StatisticsHandler.PostTrackClick(c.Request.Context(), shortUrlId)

	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error saving statistics", err)
		return
	}
	c.Status(http.StatusOK)
}

func (controller *StatisticsController) GetClickStatistics(c *gin.Context) {
	shortUrlId := c.Param("shortUrlId")
	print("SHORT URL " + shortUrlId)
	res, err := controller.AppContainer.StatisticsHandler.GetClickStatistics(c.Request.Context(), shortUrlId)

	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error getting statistics", err)
		return
	}
	c.JSON(http.StatusOK, res)
}
