package api

import (
	db "github.com/Justinecav/simplebank.git/db/sqlc"
	"github.com/gin-gonic/gin"
)

//Server serves http requests for banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

//NewServer creates a new http server and set up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	server.router = router
	return server
}

//Start runs the http server at a specific address
func (sever *Server) Start(address string) error {
	return sever.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
