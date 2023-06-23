package api

import (
	"database/sql"
	database "ecom/database/sqlc"
	"ecom/database/util"
	"ecom/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.CreateUserParams{
		ID:          util.RandomInt(1, 2000000),
		Email:       req.Email,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		CreatedAt:   time.Now(),
		CountryCode: req.CountryCode,
		FullName:    req.FullName,
		UpdatedAt:   time.Now(),
	}

	account, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) getUser(ctx *gin.Context) {
	var req models.GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req models.ListUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
