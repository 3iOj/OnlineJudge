package contest

import (
	"database/sql"
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	db "github.com/thewackyindian/3iOj/db/sqlc"
	"github.com/thewackyindian/3iOj/token"
	util "github.com/thewackyindian/3iOj/utils"
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

type createContestRequest struct {
	ContestName       string    `json:"contest_name" binding:"required"`
	StartTime         time.Time `json:"start_time" binding:"required"`
	EndTime           time.Time `json:"end_time" binding:"required"`
	Duration          int64     `json:"duration" binding:"required"`
	RegistrationStart time.Time `json:"registration_start" binding:"required"`
	RegistrationEnd   time.Time `json:"registration_end" binding:"required"`
	AnnouncementBlog  int64     `json:"announcement_blog" binding:"required"`
	EditorialBlog     int64     `json:"editorial_blog" binding:"required"`
}	
func (handler *Handler) CreateContest(ctx *gin.Context) {
	var req createContestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
		"error" : err.Error(),
		})
		return
	}
	arg := db.CreateContestTxParams{
		ContestName: req.ContestName,
		StartTime: req.StartTime,
		EndTime:  req.EndTime,
		Duration: req.Duration,      
		RegistrationStart: req.RegistrationStart, 
		RegistrationEnd: req.StartTime,  
		AnnouncementBlog: req.AnnouncementBlog,
		EditorialBlog: req.EditorialBlog,
	}
	result, err := handler.store.CreateUserTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, result.Contest)
	
}
type getContestRequest struct {
	ID int64 `uri:"id" binding:"required,num"`
}

func (handler *Handler) GetContest(ctx *gin.Context) {
	var req getContestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}
	contest, err := handler.store.GetContest(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error" : err.Error(),
		});
			return
		}
		ctx.JSON(http.StatusInternalServerError,gin.H{
				"error" : err.Error(),
		});
		return
	}

	ctx.JSON(http.StatusOK, contest)
}


//  binding logic here ?...
type listContestsRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (handler *Handler) ListContests(ctx *gin.Context) {
	var req listContestsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
		});
		return
	}
	fmt.Println(req.PageID, req.PageSize)
	if req.PageID == 0 {
		req.PageID = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 5
	}
	arg := db.ListContestsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	contests, err := handler.store.ListContests(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
				"error" : err.Error(),
		});
		return
	}
	fmt.Println(contests)

	ctx.JSON(http.StatusOK, contests)
}