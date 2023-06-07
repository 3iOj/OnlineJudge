package api

import (
	"github.com/gin-gonic/gin"
	user "github.com/thewackyindian/3iOj/api/users"
	contest "github.com/thewackyindian/3iOj/api/contests"

	db "github.com/thewackyindian/3iOj/db/sqlc"
)

//here we implement our HTTP API server
type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

    userHandler := user.NewHandler(
        // server.config,
        server.store,
        // server.tokenMaker,
		
    )

    router.POST("/users", userHandler.CreateUser)


	contestHandler := contest.NewHandler(
        // server.config,
        server.store,
        // server.tokenMaker,
		
    )

    router.POST("/contest", contestHandler.CreateContest)
	server.router = router

	return server
}


func (server *Server) Start(address string) error{
	return server.router.Run(address) 
}

// func errorResponse(err error) gin.H{
// 	return gin.H{
// 		"error" : err.Error(),
// 	}
// }