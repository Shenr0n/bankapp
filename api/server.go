package api

import (
	db "github.com/Shenr0n/bankapp/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Serves the http requests for banking services
type Server struct {
	store db.Store
	//Router to send each API request to corresponding handlers
	router *gin.Engine
}
type getIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Create a new http server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//Add routes to router

	//Account
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PUT("/accounts/:id/update/name", server.updateAccountName)
	router.DELETE("/accounts/:id/delete", server.deleteAccount)

	//Entry
	router.GET("/accounts/:id/entries", server.listAccountEntries)
	router.GET("/entries/:id", server.getEntry)

	//Transfer
	router.POST("/transfers", server.createTransfer)

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
