package contest

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/3iOj/OnlineJudge/api/middleware"
	db "github.com/3iOj/OnlineJudge/db/sqlc"
	"github.com/3iOj/OnlineJudge/token"
	util "github.com/3iOj/OnlineJudge/utils"
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

type createBlogRequest struct {
	BlogTitle   string `json:"blog_title"`
	BlogContent string `json:"blog_content"`
	Ispublish   bool   `json:"ispublish"`
}

func (handler *Handler) CreateBlog(ctx *gin.Context) {
	var req createBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(ctx)
	//ctx.mustget returns a general interface therefore we are casting it to a token.payload object
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.CreateBlogParams{
		BlogTitle:   req.BlogTitle,
		BlogContent: req.BlogContent,
		CreatedBy:   authPayload.Username,
		Ispublish:   sql.NullBool{Bool: req.Ispublish, Valid: true},
	}
	blog, err := handler.store.CreateBlog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, blog)

}

type getBlogRequest struct {
	ID int64 `uri:"id" binding:"required,num"`
}

func (handler *Handler) GetBlog(ctx *gin.Context) {
	var req getBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}
	blog, err := handler.store.GetContest(ctx, req.ID)
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

	ctx.JSON(http.StatusOK, blog)
}

type listBlogsRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (handler *Handler) ListBlogs(ctx *gin.Context) {
	var req listBlogsRequest
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
	arg := db.ListBlogsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	blogs, err := handler.store.ListBlogs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

type updateBlog struct {
	id int64 `uri:"id" binding:"required"`
}
type updateBlogRequest struct {
	BlogTitle   string `json:"blog_title"`
	BlogContent string `json:"blog_content"`
	CreatedBy   string `json:"created_by"`
	Ispublish   bool   `json:"ispublish"`
}

func (handler *Handler) UpdateBlog(ctx *gin.Context) {
	var blog updateBlog
	var req updateBlogRequest

	if err := ctx.ShouldBindUri(&blog); err != nil {
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
	toUpdateBlog, err := handler.store.GetBlog(ctx, blog.id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	if authPayload.Username != toUpdateBlog.CreatedBy {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	arg := db.UpdateBlogParams{
		BlogTitle:   sql.NullString{String: req.BlogTitle, Valid: true},
		BlogContent: sql.NullString{String: req.BlogContent, Valid: true},
		Ispublish:   sql.NullBool{Bool: req.Ispublish, Valid: true},
	}

	updatedBlog, err := handler.store.UpdateBlog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, updatedBlog)
}
