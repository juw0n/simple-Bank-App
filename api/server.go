package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/juw0n/simpleBank/db/sqlc"
)

// Server serve all the HTTP request for our banking app/service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server instance and setup all API route for all service on that server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//routes attached to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	// router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address to start listening for an API request.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
