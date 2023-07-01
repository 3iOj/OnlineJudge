package problem

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/3iOj/OnlineJudge/db/sqlc"
	"github.com/3iOj/OnlineJudge/token"
	util "github.com/3iOj/OnlineJudge/utils"
	"github.com/gin-gonic/gin"
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

type createProblemRequest struct {
	ProblemName  string `json:"problem_name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	ContestID    int64  `json:"contest_id" binding:"required"`
}

func (handler *Handler) CreateProblem(ctx *gin.Context) {
	var req createProblemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	arg := db.CreateProblemParams {
		ProblemName : req.ProblemName,
    	Description : req.Description,
		ContestID: req.ContestID,
	}

	problem, err := handler.store.CreateProblem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, problem) 
}

type getProblemRequest struct {
	ID int64 `uri:"id" binding:"required,alphanum"`
}

func (handler *Handler) GetProblem(ctx *gin.Context) {
	var req getProblemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}
	problem, err := handler.store.GetProblem(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, problem)
}
func (handler *Handler) UpdateProblem(ctx *gin.Context) {
}


func (handler *Handler) ListProblems(ctx *gin.Context) {
}
type listProblemsRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (handler *Handler) ListBlogs(ctx *gin.Context) {
	var req listProblemsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(req.PageID, req.PageSize)
	if req.PageID == 0 {
		req.PageID = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 5
	}
	arg := db.ListProblemsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	problems, err := handler.store.ListProblems(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, problems)
}



