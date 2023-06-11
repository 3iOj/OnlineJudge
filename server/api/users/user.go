package user

import (
	"database/sql"
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// "gopkg.in/guregu/null.v4"
	db "github.com/thewackyindian/3iOj/db/sqlc"
	util "github.com/thewackyindian/3iOj/utils"
)

type Handler struct {
	// config     util.Config
	store db.Store
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

type userResponse struct {
	Username      string    `json:"username"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Dob           time.Time `json:"dob"`
	Profileimg    string    `json:"profileimg"`
	Motto         string    `json:"motto"`
	Rating        int32     `json:"rating"`
	ProblemSolved int32     `json:"problem_solved"`
	AdminID       int64     `json:"admin_id"`
	IsSetter      bool      `json:"is_setter"`
	CreatedAt     time.Time `json:"created_at"`
}
type createUserRequest struct {
	Username   string    `json:"username" binding:"required,alphanum"`
	Password   string    `json:"password" binding:"required,min=8"`
	Name       string    `json:"name" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Dob        time.Time `json:"dob" binding:"required"`
	Profileimg string    `json:"profileimg"`
	Motto      string    `json:"motto"`
	IsSetter   bool      `json:"is_setter"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:      user.Username,
		Name:          user.Name,
		Email:         user.Email,
		Dob:           user.Dob,
		Profileimg:    user.Profileimg.String,
		Motto:         user.Motto.String,
		Rating:        user.Rating.Int32,
		ProblemSolved: user.ProblemSolved.Int32,
		AdminID:       user.AdminID.Int64,
		IsSetter:      user.IsSetter,
		CreatedAt:     user.CreatedAt,
	}
}
func (handler *Handler) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	arg := db.CreateUserParams{
		Name:       req.Name,
		Username:   req.Username,
		Email:      req.Email,
		Password:   hashedPassword,
		Dob:        req.Dob,
		Profileimg: sql.NullString{String: req.Profileimg, Valid: true},
		Motto:      sql.NullString{String: req.Motto, Valid: true},
	}
	user, err := handler.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	// if pqErr, ok := err.(*pq.Error); ok {
	// 	switch pqErr.Code.Name() {
	// 	case "unique_violation":
	// 		ctx.JSON(http.StatusForbidden, errorResponse(err))
	// 		return
	// 	}
	// }
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,alphanum"`
}

func (handler *Handler) GetUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}
	user, err := handler.store.GetUser(ctx, req.Username)
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
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (handler *Handler) ListUsers(ctx *gin.Context) {
	var req listUsersRequest
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
	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := handler.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	rsp := []userResponse{}
	/*
		binding all users to response style
	*/
	for i := 0; i < len(users); i++ {
		rsp = append(rsp, newUserResponse(users[i]))
	}
	ctx.JSON(http.StatusOK, rsp)
}

type updateUser struct {
	Username string `uri:"username" binding:"required,alphanum"`
}
type updateUserRequest struct {
    Password   string `json:"password"`
    Name       string `json:"name"`
    Email      string `json:"email"`
    Dob        time.Time `json:"dob"`
    Profileimg string `json:"profileimg"`
    Motto      string `json:"motto"`
    IsSetter   bool   `json:"is_setter"`
}

func (handler *Handler) UpdateUser(ctx *gin.Context) {
	var user updateUser
	var req updateUserRequest
	//incoming request must have nested json objects
	//bceause sql.nullstring requires two keys, string and valid for
	//smooth binding here -- take care of this
	if err := ctx.ShouldBindUri(&user); err != nil {
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
	hashedUpdatedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	arg := db.UpdateUserParams{
		Name:     sql.NullString{String: req.Name, Valid:true},
		Username: user.Username,
		Email:    sql.NullString{String: req.Email, Valid:true},
		Password: sql.NullString{
			String: hashedUpdatedPassword,
			Valid:  true,
		},
		Dob:        sql.NullTime{Time: req.Dob, Valid:true},
		Profileimg: sql.NullString{String: req.Profileimg, Valid: true},
		Motto:      sql.NullString{String: req.Motto, Valid: true},
	}

	updatedUser, err := handler.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rsp := newUserResponse(updatedUser)
	ctx.JSON(http.StatusOK, rsp)
}
