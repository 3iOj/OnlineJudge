package contest

import (
	"database/sql"
	"fmt"
	"net/http"

	"time"

	// "github.com/3iOj/OnlineJudge/api/middleware"
	db "github.com/3iOj/OnlineJudge/db/sqlc"
	"github.com/3iOj/OnlineJudge/token"
	util "github.com/3iOj/OnlineJudge/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	ContestName string `json:"contest_name" binding:"required"`
	Duration    int64  `json:"duration" binding:"required"`
}

	
func (handler *Handler) CreateContest(ctx *gin.Context) {//submit
	var req createContestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
		"error" : err.Error(),
		})
		return
	}
	arg := db.CreateContestParams{
		ContestName: req.ContestName,
		Duration: req.Duration,      
	}
	result, err := handler.store.CreateContest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, result)
	
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
	Ispublish      	  bool      `json:"is_publish" binding:"required"`
}

func (handler *Handler) UpdateContest(ctx *gin.Context) {
	var contest updateContest
	var req updateContestRequest

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
	// authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	// if authPayload.Username != contest.Username {
	// 	err := errors.New("account doesn't belong to the authenticated user")
	// 	ctx.JSON(http.StatusUnauthorized,gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	
	arg := db.UpdateContestTxParams{
		ContestName	: req.ContestName,
		StartTime   : req.StartTime,     
		EndTime    	: req.EndTime,       
		Duration    : req.Duration,      
		RegistrationStart : req.RegistrationStart,
		RegistrationEnd   : req.RegistrationEnd,
		AnnouncementBlog : req.AnnouncementBlog, 
		EditorialBlog   : req.EditorialBlog,  
		ContestCreators : req.ContestCreators,  
		Ispublish      : req.Ispublish,
	}

	updatedContest, err := handler.store.UpdateContestTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rsp := updatedContest
	ctx.JSON(http.StatusOK, rsp)
}

// type UpdateContestTxParams struct {
// 	ID				  int64		`json:"id" binding:"required"`
// 	ContestName       string    `json:"contest_name" binding:"required"`
// 	StartTime         time.Time `json:"start_time" binding:"required"`
// 	EndTime           time.Time `json:"end_time" binding:"required"`
// 	Duration          int64     `json:"duration" binding:"required"`
// 	RegistrationStart time.Time `json:"registration_start" binding:"required"`
// 	RegistrationEnd   time.Time `json:"registration_end" binding:"required"`
// 	AnnouncementBlog  int64     `json:"announcement_blog" binding:"required"`
// 	EditorialBlog     int64     `json:"editorial_blog" binding:"required"`
// 	ContestCreators   []string  `json:"contest_creators" binding:"required"`
// 	Ispublish         bool   `json:"ispublish"`
// }

// type UpdateContestTxResponse struct {
// 	Contest         db.Contest  `json:"contest"`
// 	ContestCreators []string `json:"contest_creators"`
// }

// func (handler *Handler) UpdateContestTx(ctx context.Context, arg UpdateContestTxParams) (UpdateContestTxResponse, error) {
// 	var rsp UpdateContestTxResponse
	
// 	err := db.SQLStore.execTx(ctx, func(q *Queries) error {
// 		var err error
// 		var announcement_blog, editorial_blog db.Blog
// 		if arg.AnnouncementBlog != 0 {
// 			announcement_blog, err = q.UpdateBlog(ctx, db.UpdateBlogParams{
// 				PublishAt: sql.NullTime{Time: time.Now(), Valid: true},
// 				ID:        arg.AnnouncementBlog,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}		
		
// 		if arg.EditorialBlog != 0 {
// 			editorial_blog, err = q.UpdateBlog(ctx,
// 			db.UpdateBlogParams{
// 				PublishAt: sql.NullTime{Time: time.Now(), Valid: true},
// 				ID:        arg.EditorialBlog,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		if arg.Ispublish == true {
// 			//checking all arguments should be NOT NULL
// 		}
// 		rsp.Contest, err = q.UpdateContest(ctx,
// 			db.UpdateContestParams{
// 				ContestName:       sql.NullString{String: arg.ContestName, Valid: true},
// 				EndTime:           sql.NullTime{Time: arg.EndTime, Valid: true},
// 				StartTime:         sql.NullTime{Time: arg.StartTime, Valid: true},
// 				Duration:          sql.NullInt64{Int64:arg.Duration, Valid: true},
// 				RegistrationStart: sql.NullTime{Time: arg.RegistrationStart, Valid: true},
// 				RegistrationEnd:   sql.NullTime{Time: arg.RegistrationEnd, Valid: true},
// 				AnnouncementBlog:  sql.NullInt64{Int64: announcement_blog.ID, Valid: true},
// 				EditorialBlog:     sql.NullInt64{Int64: editorial_blog.ID, Valid: true},
// 				UpdatedAt:         sql.NullTime{Time: arg.EndTime, Valid: true},
// 				Ispublish:         sql.NullBool{Bool: arg.Ispublish, Valid: true},
// 			})
// 		if err != nil {
// 			return err
// 		}
// 		err = q.DeleteContestCreators(ctx, rsp.Contest.ID)
// 		if err != nil {
// 			return err
// 		}
// 		contestCreators := []string{}
// 		for i := 0; i < len(arg.ContestCreators); i++ {
// 			contestCreator, err := q.AddContestCreators(ctx,
// 				db.AddContestCreatorsParams{
// 					ContestID:   arg.ID,
// 					CreatorName: arg.ContestCreators[i],
// 				})
// 			contestCreators = append(contestCreators, contestCreator.CreatorName)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		rsp.ContestCreators = contestCreators

// 		return err
// 	})
// 	return rsp, err
// }
