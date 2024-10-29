package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/registry"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/utils"
	"net/http"
)

type StatisticsController struct {
	AppContainer registry.AppContainer
}

func (controller *StatisticsController) RegisterRoutes(router gin.IRouter) {
	router.GET("/:shortUrlId", controller.GetClickStatistics)
}

func (controller *StatisticsController) GetClickStatistics(c *gin.Context) {
	shortUrlId := c.Param("shortUrlId")
	res, err := controller.AppContainer.StatisticsHandler.GetClickStatistics(c.Request.Context(), shortUrlId)

	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error getting statistics", err)
		return
	}
	c.JSON(http.StatusOK, res)
}
