package problem

import (
	db "github.com/3iOj/OnlineJudge/db/sqlc"
	"github.com/3iOj/OnlineJudge/token"
	util "github.com/3iOj/OnlineJudge/utils"
)

type Handler struct {
	config     util.Config
	store db.Store
	tokenMaker token.Maker
}
func NewHandler(
	config util.Config,
	store db.Store,
	tokenMaker token.Maker,
) *Handler {
	return &Handler{
		config,store, tokenMaker,
	}
}


