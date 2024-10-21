package api

import (
	db "github.com/bastiendmt/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// Creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
