package problem

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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

type createProblemRequest struct {
	ProblemName string `json:"problem_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ContestID   int64  `json:"contest_id" binding:"required"`
}

func (handler *Handler) CreateProblem(ctx *gin.Context) {
	var req createProblemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	arg := db.CreateProblemParams{
		ProblemName: req.ProblemName,
		Description: req.Description,
		ContestID:   req.ContestID,
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

type updateProblem struct {
	problemid int64 `uri:"id" binding:"required,alphanum"`
}
type updateUserRequest struct {
	Problemname   string `json:"problem_name"`
	Description   string `json:"description"`
	SampleInput   string `json:"sample_input"`
	SampleOutput  string `json:"sample_output"`
	IdealSolution string `json:"ideal_solution"`
	TimeLimit     int32  `json:"time_limit"`
	MemoryLimit   int32  `json:"memory_limit"`
	CodeSize      int32  `json:"code_size"`
	Rating        int32  `json:"rating"`
}

func (handler *Handler) UpdateProblem(ctx *gin.Context) {
	var problem updateProblem
	var req updateUserRequest

	if err := ctx.ShouldBindUri(&problem); err != nil {
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
	arg := db.UpdateProblemParams{
		ID:            problem.problemid,
		ProblemName:   sql.NullString{String: req.Problemname, Valid: true},
		Description:   sql.NullString{String: req.Description, Valid: true},
		SampleInput:   sql.NullString{String: req.SampleInput, Valid: true},
		SampleOutput:  sql.NullString{String: req.SampleOutput, Valid: true},
		IdealSolution: sql.NullString{String: req.IdealSolution, Valid: true},
		TimeLimit:     sql.NullInt32{Int32: req.TimeLimit, Valid: true},
		MemoryLimit:   sql.NullInt32{Int32: req.MemoryLimit, Valid: true},
		CodeSize:      sql.NullInt32{Int32: req.CodeSize, Valid: true},
		Rating:        sql.NullInt32{Int32: req.Rating, Valid: true},
	}
	updatedProblem, err := handler.store.UpdateProblem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotModified, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedProblem)
}

type submitTestCases struct {
	problemid int64 `uri:"id" binding:"required,alphanum"`
}

type submitTestCasesRequest struct {
	TestCases *multipart.FileHeader `form:"testcases" binding:"required"`
}

func (handler *Handler) SubmitTestCases(ctx *gin.Context) {
	var req submitTestCasesRequest
	var problem submitTestCases

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindUri(&problem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tempDir, err := ioutil.TempDir("", "uploads")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer os.RemoveAll(tempDir)

	uploadedFilePath := filepath.Join(tempDir, req.TestCases.Filename)
	if err := ctx.SaveUploadedFile(req.TestCases, uploadedFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	go func() {
		util.Unzip(uploadedFilePath, tempDir)
		util.UploadFile(tempDir, problem.problemid)
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file uploaded successfully",
	})
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
