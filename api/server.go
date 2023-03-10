package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/juw0n/simpleBank/db/sqlc"
)

// Server serve all the HTTP request fro our banking app
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.getAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a sspecific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
