package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
	"github.com/jesusbibieca/url-shortener/environment"
	"github.com/jesusbibieca/url-shortener/redis_store"
	"github.com/jesusbibieca/url-shortener/shortener"
)

type ShortUrlCreateRequest struct {
	Url    string `json:"url" binding:"required"`
	UserId string `json:"userId" binding:"required"`
}

func (server *Server) getShortUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	url := redis_store.RetrieveInitialUrl(shortUrl)

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (server *Server) createShortUrl(ctx *gin.Context) {
	config, err := environment.LoadConfig("../")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var shortUrlRequest ShortUrlCreateRequest
	if err := ctx.ShouldBindJSON(&shortUrlRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	shortUrl := shortener.GenerateShortLink(shortUrlRequest.Url, shortUrlRequest.UserId)
	redis_store.SaveUrlMapping(shortUrl, shortUrlRequest.Url, shortUrlRequest.UserId)

	ctx.JSON(http.StatusOK, gin.H{
		// TODO: refactor this
		"shortUrl": "http://" + config.AppAddress + shortUrl,
	})
}

type ShortUrlUpdateRequest struct {
	OriginalUrl string `json:"originalUrl" binding:"required"`
}

// router.PATCH("/r/:shortUrl", server.updateShortUrl)
func (server *Server) updateShortUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")

	var shortUrlRequest ShortUrlUpdateRequest
	if err := ctx.ShouldBindJSON(&shortUrlRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	url := pgtype.Text{
		String: shortUrl,
		Valid:  true,
	}

	// get url from db
	_, err := server.store.GetShortUrl(ctx, url)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// TODO: verify if the user is the owner of the url

	// update url in db
	updatedUrl, err := server.store.UpdateShortUrl(ctx, db.UpdateShortUrlParams{
		OriginalUrl: shortUrlRequest.OriginalUrl,
		ShortUrl:    url,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, updatedUrl)
}

// router.DELETE("/r/:shortUrl", server.deleteShortUrl)
func (server *Server) deleteShortUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")

	url := pgtype.Text{
		String: shortUrl,
		Valid:  true,
	}

	// get url from db
	_, err := server.store.GetShortUrl(ctx, url)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// TODO: verify if the user is the owner of the url

	// delete url from db
	err = server.store.DeleteShortUrl(ctx, url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "url deleted",
	})
}

func (server *Server) getPagedUrls(ctx *gin.Context) {
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", fmt.Sprint(db.DefaultLimit)), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit32 := int32(limit)

	if limit > db.MaxLimit {
		limit32 = db.MaxLimit
	}

	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", "0"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset32 := int32(offset)

	urls, err := server.store.GetPagedUrls(ctx, db.GetPagedUrlsParams{
		Limit:  limit32,
		Offset: offset32,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, urls)

}
