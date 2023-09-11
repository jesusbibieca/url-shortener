package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func (server *Server) getUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id32 := int32(id)

	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid user id")))
		return
	}

	user, err := server.store.GetUserById(ctx, id32)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) deleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id32 := int32(id)

	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid user id")))
		return
	}

	err = server.store.DeleteUser(ctx, id32)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (server *Server) getPagedUsers(ctx *gin.Context) {
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

	users, err := server.store.GetPagedUsers(ctx, db.GetPagedUsersParams{
		Limit:  limit32,
		Offset: offset32,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
