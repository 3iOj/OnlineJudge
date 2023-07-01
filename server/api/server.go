package api

import (
	"fmt"

	"github.com/3iOj/OnlineJudge/api/admin"
	blog "github.com/3iOj/OnlineJudge/api/blogs"
	problem "github.com/3iOj/OnlineJudge/api/problems"
	contest "github.com/3iOj/OnlineJudge/api/contests"
	"github.com/3iOj/OnlineJudge/api/middleware"
	user "github.com/3iOj/OnlineJudge/api/users"
	db "github.com/3iOj/OnlineJudge/db/sqlc"
	"github.com/3iOj/OnlineJudge/token"
	util "github.com/3iOj/OnlineJudge/utils"
	"github.com/gin-gonic/gin"
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
	
	server.setupRouter()
	return server, err
}


func(server *Server)  setupRouter() {
router := gin.Default()

	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))
	adminHandler := admin.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)
	authRoutes.POST("/admin/register", adminHandler.CreateAdmin)
    userHandler := user.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)
	
	router.POST("/users/register", userHandler.CreateUser)
    router.POST("/users/login", userHandler.LoginUser)
    // router.GET("/users", userHandler.ListUsers)
    router.GET("/users/:username", userHandler.GetUser)//profile page
	
	
    
	// authRoutes.POST("/admin/register", adminHandler.CreateAdmin)
	authRoutes.PUT("/users/:username/", userHandler.UpdateUser)
	
	
	
    contestHandler := contest.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)
	
	router.GET("/contests/:username",contestHandler.GetContest)
	router.GET("/contests", contestHandler.ListContests)
	authRoutes.POST("/contests/create", contestHandler.CreateContest)
	authRoutes.GET("/contest/:id", contestHandler.GetContest)
	authRoutes.PUT("/contests/edit/:id", contestHandler.UpdateContest)

	blogHandler := blog.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)

	authRoutes.POST("/blogs", blogHandler.CreateBlog)
	router.GET("/blogs", blogHandler.ListBlogs)
	router.GET("/blogs/:id", blogHandler.GetBlog)
	authRoutes.PUT("/blogs/:id",blogHandler.UpdateBlog)

	

	problemHandler := problem.NewHandler(
		server.config,
		server.store,
		server.tokenMaker,
	)
	authRoutes.POST("/problems", problemHandler.CreateProblem)
	router.GET("/problems", problemHandler.ListProblems)
	router.GET("/problems/:id", problemHandler.GetProblem)
	authRoutes.PUT("/problems/:id",problemHandler.UpdateProblem)
	server.router = router
}


func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H{
	return gin.H{
		"error" : err.Error(),
	}
}
