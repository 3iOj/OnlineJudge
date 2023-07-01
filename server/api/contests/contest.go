package contest

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"time"

	// "github.com/3iOj/OnlineJudge/api/middleware"
	"github.com/3iOj/OnlineJudge/api/middleware"
	db "github.com/3iOj/OnlineJudge/db/sqlc"
	"github.com/3iOj/OnlineJudge/token"
	util "github.com/3iOj/OnlineJudge/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

type createContestRequest struct {
	ContestName string `json:"contest_name" binding:"required"`
	Duration    int64  `json:"duration" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
}
type contestResponse struct {
	Contest         db.Contest `json:"contest"`
	ContestCreators []string   `json:"contest_creators"`
}

func (handler *Handler) CreateContest(ctx *gin.Context) { //submit
	var req createContestRequest
	var rsp contestResponse
	var err error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	arg := db.CreateContestParams{
		ContestName: req.ContestName,
		Duration:    req.Duration,
	}
	rsp.Contest, err = handler.store.CreateContest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	contestCreators := []string{}
	contestCreator, err := handler.store.AddContestCreators(ctx,
		db.AddContestCreatorsParams{
			ContestID:   rsp.Contest.ID,
			CreatorName: req.CreatedBy,
		})
	contestCreators = append(contestCreators, contestCreator.CreatorName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rsp.ContestCreators = contestCreators
	ctx.JSON(http.StatusOK, rsp)

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
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, contest)
}

type listContestsRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (handler *Handler) ListContests(ctx *gin.Context) {
	var req listContestsRequest
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
	arg := db.ListContestsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	contests, err := handler.store.ListContests(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(contests)

	ctx.JSON(http.StatusOK, contests)
}

type updateContest struct {
	ID int64 `uri:"id" binding:"required,alphanum"`
}
type updateContestRequest struct {
	ContestName       string    `json:"contest_name" binding:"required"`
	StartTime         time.Time `json:"start_time" binding:"required"`
	EndTime           time.Time `json:"end_time" binding:"required"`
	Duration          int64     `json:"duration" binding:"required"`
	RegistrationStart time.Time `json:"registration_start" binding:"required"`
	RegistrationEnd   time.Time `json:"registration_end" binding:"required"`
	AnnouncementBlog  int64     `json:"announcement_blog" binding:"required"`
	EditorialBlog     int64     `json:"editorial_blog" binding:"required"`
	ContestCreators   []string  `json:"contest_creators" binding:"required"`
	Ispublish         bool      `json:"is_publish" binding:"required"`
}

func (handler *Handler) UpdateContest(ctx *gin.Context) {
	var contest updateContest
	var req updateContestRequest
	var rsp contestResponse
	// var err error
	if err := ctx.ShouldBindUri(&contest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	
	
	validContestCreators, err := handler.store.GetContestCreators(ctx, contest.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var ok bool
	for i := 0; i < len(validContestCreators); i++ {
		if authPayload.Username == validContestCreators[i] {
			ok = true
			break
		}
	}
	if !ok {
		err = errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	arg := db.UpdateContestParams{
		ContestName:       sql.NullString{String: req.ContestName, Valid: true},
		EndTime:           sql.NullTime{Time: req.EndTime, Valid: true},
		StartTime:         sql.NullTime{Time: req.StartTime, Valid: true},
		Duration:          sql.NullInt64{Int64: req.Duration, Valid: true},
		RegistrationStart: sql.NullTime{Time: req.RegistrationStart, Valid: true},
		RegistrationEnd:   sql.NullTime{Time: req.RegistrationEnd, Valid: true},
		AnnouncementBlog:  sql.NullInt64{Int64: req.AnnouncementBlog, Valid: true},
		EditorialBlog:     sql.NullInt64{Int64: req.EditorialBlog, Valid: true},
		UpdatedAt:         sql.NullTime{Time: req.EndTime, Valid: true},
		Ispublish:         sql.NullBool{Bool: req.Ispublish, Valid: true},
	}
	rsp.Contest, err = handler.store.UpdateContest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = handler.store.DeleteContestCreators(ctx, contest.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	contestCreators := []string{}
	for i := 0; i < len(req.ContestCreators); i++ {
		contestCreator, err := handler.store.AddContestCreators(ctx,
			db.AddContestCreatorsParams{
				ContestID:   arg.ID,
				CreatorName: req.ContestCreators[i],
			})
		contestCreators = append(contestCreators, contestCreator.CreatorName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	rsp.ContestCreators = contestCreators
	ctx.JSON(http.StatusOK, rsp)
}
