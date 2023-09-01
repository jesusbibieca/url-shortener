package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Service is up and running 🚀"})
}
