package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/registry"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/utils"
	"net/http"
)

type UrlController struct {
	AppContainer registry.AppContainer
}

func (controller *UrlController) RegisterRoutes(router gin.IRouter) {
	router.POST("/shorten", controller.PostShorten)
	router.GET("/:shortUrlId", controller.GetLongUrl)
}

func (controller *UrlController) GetLongUrl(c *gin.Context) {
	var err error
	shortUrlId := c.Param("shortUrlId")

	longUrl, err := controller.AppContainer.UrlHandler.GetLongUrl(c.Request.Context(), shortUrlId)
	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error on getting url_manage", err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, longUrl)
}

func (controller *UrlController) PostShorten(c *gin.Context) {
	payload := domain.ShortenUrlRequest{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusBadRequest, "error on bindJson", err)
		return
	}
	response, err := controller.AppContainer.UrlHandler.ShortenUrl(c.Request.Context(), payload)

	if err != nil {
		utils.AbortWithStatusCode(c, http.StatusInternalServerError, "error on get ", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
