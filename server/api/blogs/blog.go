package contest

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	db "github.com/thewackyindian/3iOj/db/sqlc"
)
type Handler struct {
    // config     util.Config
    store      db.Store
    // tokenMaker token.Maker
	
}

func NewHandler(
    // config util.Config,
    store db.Store,
    // tokenMaker token.Maker,
) *Handler {
    return &Handler{
         store, 
    }
}

type createBlogRequest struct {
	BlogTitle   string `json:"blog_title"`
	BlogContent string `json:"blog_content"`
	CreatedBy   string  `json:"created_by"`
	PublishAt   time.Time `json:"publish_at"`
}	
func (handler *Handler) CreateBlog(ctx *gin.Context) {
	var req createBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
		"error" : err.Error(),
		})
		return
	}

	arg := db.CreateBlogParams{
		BlogTitle: req.BlogTitle,
	    BlogContent:req.BlogContent,
	    CreatedBy: req.CreatedBy,
		PublishAt: req.PublishAt,
	}
	blog, err := handler.store.CreateBlog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
	}

	// Perform the redirection to the created blog page
	blogID := blog.ID
	redirectURL := fmt.Sprintf("/blogs/%d", blogID)
	ctx.Redirect(http.StatusMovedPermanently, redirectURL)

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
				"error" : err.Error(),
		});
			return
		}
		ctx.JSON(http.StatusInternalServerError,gin.H{
				"error" : err.Error(),
		});
		return
	}

	ctx.JSON(http.StatusOK, blog)
}


// binding logic here ?...
type listBlogsRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (handler *Handler) ListBlogs(ctx *gin.Context) {
	var req listBlogsRequest
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
	arg := db.ListBlogsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	blogs, err := handler.store.ListBlogs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
				"error" : err.Error(),
		});
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}