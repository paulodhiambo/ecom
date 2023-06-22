package api

import (
	database "ecom/database/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *database.Store
	router *gin.Engine
}

func NewServer(store *database.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
