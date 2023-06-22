package api

import (
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

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
