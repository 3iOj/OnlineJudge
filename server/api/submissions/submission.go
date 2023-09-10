package submission

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	db "github.com/3iOj/OnlineJudge/db/sqlc"
// 	"github.com/3iOj/OnlineJudge/token"
// 	util "github.com/3iOj/OnlineJudge/utils"
// 	"github.com/gin-gonic/gin"
// )

// type Handler struct {
// 	config     util.Config
// 	store      db.Store
// 	tokenMaker token.Maker
// }

// func NewHandler(
// 	config util.Config,
// 	store db.Store,
// 	tokenMaker token.Maker,
// ) *Handler {
// 	return &Handler{
// 		config, store, tokenMaker,
// 	}
// }

// type createSubmissionRequest struct {
// 	ProblemID int64  `json:"problem_id"  binding:"required"`
// 	Username  string `json:"username"  binding:"required"`
// 	UserID    int64  `json:"user_id"  binding:"required"`
// 	ContestID int64  `json:"contest_id"  binding:"required"`
// 	Language  string `json:"language"  binding:"required"`
// 	Code      string `json:"code"  binding:"required"`
// 	SubmittedAt int64 `json:"submitted_at" binding:"required"`
// }

// func (handler *Handler) CreateSubmission(ctx *gin.Context) {
// 	var req createSubmissionRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	// temp judge
// 	arr := [4]string{"AC", "WA", "TLE"}
// 	j := util.RandomInt(0, 10)
// 	if j >= 3 {
// 		j = 0
// 	}
// 	score := 0
// 	if j == 0 {
// 		score = 100
// 	}

// 	exec_time := util.RandomInt(20, 1000)
// 	memory_consume := util.RandomInt(100, 2000)
// 	if j == 2 {
// 		exec_time = 5000
// 	}
// 	time.Sleep(time.Duration(exec_time) + 1000)
// 	arg := db.CreateSubmissionParams{
// 		ProblemID:      req.ProblemID,
// 		Username:       req.Username,
// 		UserID:         req.UserID,
// 		ContestID:      req.ContestID,
// 		Language:       req.Language,
// 		Code:           req.Code,
// 		ExecTime:       sql.NullInt32{Int32: int32(exec_time), Valid: true},
// 		Verdict:        sql.NullString{String: arr[j], Valid: true},
// 		Score:          sql.NullInt32{Int32: int32(score), Valid: true},
// 		MemoryConsumed: sql.NullInt32{Int32: int32(memory_consume), Valid: true},
// 	}

// 	Submission, err := handler.store.CreateSubmission(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 	}

	
// 	ctx.JSON(http.StatusOK, Submission)

// }

// type getSubmissionRequest struct {
// 	ID int64 `uri:"id" binding:"required,alphanum"`
// }

// func (handler *Handler) GetSubmission(ctx *gin.Context) {
// 	var req getSubmissionRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, (err))
// 		return
// 	}
// 	Submission, err := handler.store.GetSubmission(ctx, req.ID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, Submission)
// }
// func (handler *Handler) UpdateSubmission(ctx *gin.Context) {

// }

// func (handler *Handler) ListSubmissions(ctx *gin.Context) {

// }

// type listSubmissionsRequest struct {
// 	PageID   int32 `form:"page_id"`
// 	PageSize int32 `form:"page_size"`
// }

// func (handler *Handler) ListBlogs(ctx *gin.Context) {
// 	var req listSubmissionsRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	fmt.Println(req.PageID, req.PageSize)
// 	if req.PageID == 0 {
// 		req.PageID = 1
// 	}
// 	if req.PageSize == 0 {
// 		req.PageSize = 5
// 	}
// 	arg := db.ListSubmissionsParams{
// 		Limit:  req.PageSize,
// 		Offset: (req.PageID - 1) * req.PageSize,
// 	}

// 	Submissions, err := handler.store.ListSubmissions(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, Submissions)
// }