package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/registry"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/utils"
	"net/http"
)

type UrlManageController struct {
	AppContainer registry.AppContainer
}

func (controller *UrlManageController) RegisterRoutes(router gin.IRouter) {
	router.PATCH("/:shortUrlId ", controller.PatchUrl)
	router.POST("/:shortUrlId/enable ", controller.PostEnable)
	router.POST("/:shortUrlId/disable ", controller.PostDisable)
}

func (controller *UrlManageController) PatchUrl(c *gin.Context) {
	payload := domain.PatchUrlRequest{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusBadRequest, "error on bindJson", err)
		return
	}
	shortUrlId := c.Param("shortUrlId")
	payload.ShortUrlId = shortUrlId

	response, err := controller.AppContainer.UrlManageHandler.UpdateUrl(c.Request.Context(), payload)

	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error updating url", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller *UrlManageController) PostEnable(c *gin.Context) {
	shortUrlId := c.Param("shortUrlId")

	err := controller.AppContainer.UrlManageHandler.EnableUrl(c.Request.Context(), shortUrlId)

	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error enabling url", err)
		return
	}
	c.Status(http.StatusOK)
}

func (controller *UrlManageController) PostDisable(c *gin.Context) {
	shortUrlId := c.Param("shortUrlId")

	err := controller.AppContainer.UrlManageHandler.DisableUrl(c.Request.Context(), shortUrlId)

	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error disabling url", err)
		return
	}
	c.Status(http.StatusOK)
}
