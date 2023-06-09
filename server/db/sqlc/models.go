// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Blog struct {
	ID          int64         `json:"id"`
	BlogTitle   string        `json:"blog_title"`
	BlogContent string        `json:"blog_content"`
	CreatedBy   string        `json:"created_by"`
	CreatedAt   time.Time     `json:"created_at"`
	Ispublish   sql.NullBool  `json:"ispublish"`
	VotesCount  sql.NullInt32 `json:"votes_count"`
}

type BlogComment struct {
	ID           int64     `json:"id"`
	BlogID       int64     `json:"blog_id"`
	Message      string    `json:"message"`
	CommentedBy  string    `json:"commented_by"`
	ChildComment int64     `json:"child_comment"`
	CreatedAt    time.Time `json:"created_at"`
}

type BlogLike struct {
	BlogID    int64     `json:"blog_id"`
	ActionBy  string    `json:"action_by"`
	IsLiked   bool      `json:"is_liked"`
	CreatedAt time.Time `json:"created_at"`
}

type BlogTag struct {
	BlogID int64  `json:"blog_id"`
	Tag    string `json:"tag"`
}

type Community struct {
	ID             int64     `json:"id"`
	CommunityName  string    `json:"community_name"`
	CommunityAdmin string    `json:"community_admin"`
	CreatedAt      time.Time `json:"created_at"`
}

type CommunityUser struct {
	ID          int64  `json:"id"`
	CommunityID int64  `json:"community_id"`
	Username    string `json:"username"`
}

type Contest struct {
	ID          int64        `json:"id"`
	ContestName string       `json:"contest_name"`
	StartTime   sql.NullTime `json:"start_time"`
	// must be greater than start time
	EndTime sql.NullTime `json:"end_time"`
	// must be equal to difference between end time and start time
	Duration          int64        `json:"duration"`
	RegistrationStart sql.NullTime `json:"registration_start"`
	// must be greater than registration_start
	RegistrationEnd sql.NullTime `json:"registration_end"`
	// should be created automatically
	AnnouncementBlog sql.NullInt64 `json:"announcement_blog"`
	// should be created automatically
	EditorialBlog sql.NullInt64 `json:"editorial_blog"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     sql.NullTime  `json:"updated_at"`
	Ispublish     sql.NullBool  `json:"ispublish"`
}

type ContestCreator struct {
	ContestID   int64     `json:"contest_id"`
	CreatorName string    `json:"creator_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type ContestRegistered struct {
	ContestID int64     `json:"contest_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type ProbTag struct {
	ProblemID int64  `json:"problem_id"`
	Tag       string `json:"tag"`
}

type Problem struct {
	ID           int64          `json:"id"`
	ProblemName  string         `json:"problem_name"`
	Description  string         `json:"description"`
	SampleInput  sql.NullString `json:"sample_input"`
	SampleOutput sql.NullString `json:"sample_output"`
	// to generate the output files in problemTests
	IdealSolution sql.NullString `json:"ideal_solution"`
	// should be in seconds
	TimeLimit sql.NullInt32 `json:"time_limit"`
	// should be in MB
	MemoryLimit sql.NullInt32 `json:"memory_limit"`
	// should be in KB
	CodeSize  sql.NullInt32 `json:"code_size"`
	Rating    sql.NullInt32 `json:"rating"`
	CreatedAt time.Time     `json:"created_at"`
	ContestID int64         `json:"contest_id"`
}

type ProblemCreator struct {
	ProblemID int64     `json:"problem_id"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type Request struct {
	Username        string         `json:"username"`
	AdminName       string         `json:"admin_name"`
	PermissionAsked sql.NullString `json:"permission_asked"`
	CurrentStatus   sql.NullString `json:"current_status"`
	CreatedAt       time.Time      `json:"created_at"`
}

type Session struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	RefreshToken  string         `json:"refresh_token"`
	Profileimg    sql.NullString `json:"profileimg"`
	Motto         sql.NullString `json:"motto"`
	CreatedAt     time.Time      `json:"created_at"`
	Dob           time.Time      `json:"dob"`
	Rating        sql.NullInt32  `json:"rating"`
	ProblemSolved sql.NullInt32  `json:"problem_solved"`
	AdminID       sql.NullInt64  `json:"admin_id"`
	IsSetter      bool           `json:"is_setter"`
}

type Submission struct {
	ID int64 `json:"id"`
	// fetch solutions based on this time during a contest
	SubmittedAt time.Time `json:"submitted_at"`
	// will come handy in contest score calculation
	ProblemID int64  `json:"problem_id"`
	Username  string `json:"username"`
	// will come handy in contest score calculation
	UserID int64 `json:"user_id"`
	// will come handy in contest score calculation
	ContestID int64  `json:"contest_id"`
	Language  string `json:"language"`
	Verdict   string `json:"verdict"`
	Code      string `json:"code"`
	// should be in seconds
	ExecTime int32 `json:"exec_time"`
	// should be in MB
	MemoryConsumed int32 `json:"memory_consumed"`
	// must be calculated in application logic
	Score int32 `json:"score"`
}

type SubmissionTest struct {
	ID           int64         `json:"id"`
	SubmissionID sql.NullInt64 `json:"submission_id"`
	Input        string        `json:"input"`
	Output       string        `json:"output"`
}

type User struct {
	ID            int64          `json:"id"`
	Name          string         `json:"name"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	Profileimg    sql.NullString `json:"profileimg"`
	Motto         sql.NullString `json:"motto"`
	CreatedAt     time.Time      `json:"created_at"`
	Dob           time.Time      `json:"dob"`
	Rating        sql.NullInt32  `json:"rating"`
	ProblemSolved sql.NullInt32  `json:"problem_solved"`
	AdminID       sql.NullInt64  `json:"admin_id"`
	IsSetter      bool           `json:"is_setter"`
}
