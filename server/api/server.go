package api

import (
	"github.com/gin-gonic/gin"
	user "github.com/thewackyindian/3iOj/api/users"
	contest "github.com/thewackyindian/3iOj/api/contests"
	blog "github.com/thewackyindian/3iOj/api/blogs"

	db "github.com/thewackyindian/3iOj/db/sqlc"
)

//here we implement our HTTP API server
type Server struct {
	store db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

    userHandler := user.NewHandler(
        // server.config,
        server.store,
        // server.tokenMaker,
		
    )

    router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.ListUsers)
	router.GET("/users/:username", userHandler.GetUser)

	contestHandler := contest.NewHandler(
        // server.config,
        server.store,
        // server.tokenMaker,	
    )

    router.POST("/contests", contestHandler.CreateContest)
    router.GET("/contests", contestHandler.ListContests)
    router.GET("/contests/:id", contestHandler.GetContest)


	blogHandler := blog.NewHandler(
        // server.config,
        server.store,
        // server.tokenMaker,
    )


    router.POST("/blogs", blogHandler.CreateBlog)
    router.GET("/blogs", blogHandler.ListBlogs)
    router.GET("/blogs/:id", blogHandler.GetBlog)
	// router.PUT("/blogs/:id",blogHandler.Updateblog)

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