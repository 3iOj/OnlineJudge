package contest

import (
	// "database/sql"
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	db "github.com/thewackyindian/3iOj/db/sqlc"
	// util "github.com/thewackyindian/3iOj/utils"
)
type Handler struct {
    // config     util.Config
    store      *db.Store
    // tokenMaker token.Maker
	
}

func NewHandler(
    // config util.Config,
    store *db.Store,
    // tokenMaker token.Maker,
) *Handler {
    return &Handler{
         store, 
    }
}





// type getUserRequest struct {
// 	Username string `uri:"username" binding:"required,alphanum"`
// }

// func (handler *Handler) GetUser(ctx *gin.Context) {
// 	var req getUserRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, (err))
// 		return
// 	}
// 	user, err := handler.store.GetUser(ctx, req.Username)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"error" : err.Error(),
// 		});
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError,gin.H{
// 				"error" : err.Error(),
// 		});
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, user)
// }

// // binding logic here ?...
// type listUsersRequest struct {
// 	PageID   int32 `form:"page_id"`
// 	PageSize int32 `form:"page_size"`
// }

// func (handler *Handler) ListUsers(ctx *gin.Context) {
// 	var req listUsersRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error" : err.Error(),
// 		});
// 		return
// 	}
// 	fmt.Println(req.PageID, req.PageSize)
// 	if req.PageID == 0 {
// 		req.PageID = 1
// 	}
// 	if req.PageSize == 0 {
// 		req.PageSize = 5
// 	}
// 	arg := db.ListUsersParams{
// 		Limit:  req.PageSize,
// 		Offset: (req.PageID - 1) * req.PageSize,
// 	}

// 	accounts, err := handler.store.ListUsers(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error" : err.Error(),
// 		});
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, accounts)
// }

type createContestRequest struct {
	ContestName       string    `json:"contest_name"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	Duration          int64     `json:"duration"`
	RegistrationStart time.Time `json:"registration_start"`
	RegistrationEnd   time.Time `json:"registration_end"`
	AnnouncementBlog  int64     `json:"announcement_blog"`
	EditorialBlog     int64     `json:"editorial_blog"`
}	
func (handler *Handler) CreateContest(ctx *gin.Context) {
	var req createContestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
		"error" : err.Error(),
		})
		return
	}
	//end time , blogs logic

	arg := db.CreateContestParams{
		ContestName: req.ContestName,
		StartTime: req.StartTime,
		EndTime:  req.EndTime,
		Duration: req.Duration,      
		RegistrationStart: req.RegistrationStart, 
		RegistrationEnd: req.StartTime,  
		AnnouncementBlog: req.AnnouncementBlog,
		EditorialBlog: req.EditorialBlog,
	}
	contest, err := handler.store.CreateContest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
	}
	// Perform the redirection to the created contest page
	contestID := contest.ID
	redirectURL := fmt.Sprintf("/contest/%d", contestID)
	ctx.Redirect(http.StatusMovedPermanently, redirectURL)

	// ctx.JSON(http.StatusOK, contest)--?
	
}
