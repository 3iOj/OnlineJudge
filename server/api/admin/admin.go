package admin

import (
	"net/http"

	db "github.com/3iOj/OnlineJudge/db/sqlc"
	"github.com/3iOj/OnlineJudge/token"
	util "github.com/3iOj/OnlineJudge/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewHandler(
	config util.Config,
	store db.Store,
	tokenMaker token.Maker,
) *Handler {
	return &Handler{
		config, store, tokenMaker,
	}
}


type createAdminRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (handler *Handler) CreateAdmin(ctx *gin.Context) {
	var req createAdminRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}	
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	arg := db.CreateAdminParams{
		Name:       req.Name,
		Username:   req.Username,
		Email:      req.Email,
		Password:   hashedPassword,
	}
	admin, err := handler.store.CreateAdmin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, admin)
}



