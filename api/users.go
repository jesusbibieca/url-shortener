package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
	"github.com/jesusbibieca/url-shortener/helpers"
)

type CreateUserReq struct {
	Username string `json:"username" validate:"required,alphanum,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateUserResp struct {
	ID       int32              `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	CreateAt pgtype.Timestamptz `json:"createdAt"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userArgs := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, userArgs)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusConflict, errorResponse(err))
			return
		}
	}

	rsp := CreateUserResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		CreateAt: user.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, rsp)

}
