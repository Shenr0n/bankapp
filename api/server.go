package api

import (
	db "github.com/Shenr0n/bankapp/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Serves the http requests for banking services
type Server struct {
	store *db.Store
	//Router to send each API request to corresponding handlers
	router *gin.Engine
}

// Create a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//Add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PUT("/accounts/:id/update/name", server.updateAccountName)
	router.DELETE("/accounts/:id/delete", server.deleteAccount)

	server.router = router
	return server
}

// Run the http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// gin.H object is a map[string]interface{} object
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
