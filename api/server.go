package api

import (
	"fmt"

	db "github.com/Shenr0n/bankapp/db/sqlc"
	"github.com/Shenr0n/bankapp/token"
	"github.com/Shenr0n/bankapp/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Serves the http requests for banking services
type Server struct {
	config util.Config
	store  db.Store
	//Router to send each API request to corresponding handlers
	router *gin.Engine

	tokenMaker token.Maker
}
type getIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Create a new http server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("NewPasetoMaker failed: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

// Run the http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Add routes to router
func (server *Server) setupRouter() {
	router := gin.Default()
	//User
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	//router.GET("/users/:username", server.getUser

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
}

// gin.H object is a map[string]interface{} object
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
