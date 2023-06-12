package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	blog "github.com/thewackyindian/3iOj/api/blogs"
	contest "github.com/thewackyindian/3iOj/api/contests"
	user "github.com/thewackyindian/3iOj/api/users"

	auth "github.com/thewackyindian/3iOj/api/middleware"
	"github.com/thewackyindian/3iOj/token"
	util "github.com/thewackyindian/3iOj/utils"

	db "github.com/thewackyindian/3iOj/db/sqlc"
)

// here we implement our HTTP API server
type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err) //%w format specifier for original error
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

    userHandler := user.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)
	router.POST("/users/register", userHandler.CreateUser)
    router.POST("/users/login", userHandler.LoginUser)
    router.GET("/users", userHandler.ListUsers)
    router.GET("/users/:username", userHandler.GetUser) //thinking....

    authRoutes := router.Group("/").Use(auth.AuthMiddleware(server.tokenMaker))
	
	authRoutes.PUT("/users/:username/edit", userHandler.UpdateUser)
	authRoutes.GET("/users/:username/blogs", userHandler.UpdateUser)

	
    contestHandler := contest.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)
	
	router.GET("/contests", contestHandler.ListContests)
	authRoutes.POST("/contests", contestHandler.CreateContest)
	authRoutes.GET("/contest/:id", contestHandler.GetContest)

	blogHandler := blog.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)

	authRoutes.POST("/blogs", blogHandler.CreateBlog)
	authRoutes.GET("/blogs", blogHandler.ListBlogs)
	authRoutes.GET("/blogs/:id", blogHandler.GetBlog)
	// router.PUT("/blogs/:id",blogHandler.Updateblog)

	server.router = router

	return server, err
}


// func(server *Server)  setupRouter() {

// }


func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H{
	return gin.H{
		"error" : err.Error(),
	}
}
