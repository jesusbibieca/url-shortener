package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesusbibieca/url-shortener/environment"
	"github.com/jesusbibieca/url-shortener/shortener"
	"github.com/jesusbibieca/url-shortener/store"
)

type ShortUrlCreateRequest struct {
	Url    string `json:"url" binding:"required"`
	UserId string `json:"userId" binding:"required"`
}

func (server *Server) getShortUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	url := store.RetrieveInitialUrl(shortUrl)

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (server *Server) createShortUrl(ctx *gin.Context) {
	config, err := environment.LoadConfig()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var shortUrlRequest ShortUrlCreateRequest
	if err := ctx.ShouldBindJSON(&shortUrlRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	shortUrl := shortener.GenerateShortLink(shortUrlRequest.Url, shortUrlRequest.UserId)
	store.SaveUrlMapping(shortUrl, shortUrlRequest.Url, shortUrlRequest.UserId)

	ctx.JSON(http.StatusOK, gin.H{
		// refactor this
		"shortUrl": "http://" + config.AppAddress + shortUrl,
	})
}
