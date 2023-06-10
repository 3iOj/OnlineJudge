// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"
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
	PublishAt   time.Time     `json:"publish_at"`
	VotesCount  sql.NullInt32 `json:"votes_count"`
}

type BlogComment struct {
	ID           int64          `json:"id"`
	BlogID       int64          `json:"blog_id"`
	Message      sql.NullString `json:"message"`
	CommentedBy  int64          `json:"commented_by"`
	ChildComment int64          `json:"child_comment"`
	CreatedAt    time.Time      `json:"created_at"`
}

type BlogLike struct {
	BlogID    int64     `json:"blog_id"`
	ActionBy  int64     `json:"action_by"`
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
	CommunityAdmin int64     `json:"community_admin"`
	CreatedAt      time.Time `json:"created_at"`
}

type CommunityUser struct {
	ID          int64 `json:"id"`
	CommunityID int64 `json:"community_id"`
	UserID      int64 `json:"user_id"`
}

type Contest struct {
	ID          int64     `json:"id"`
	ContestName string    `json:"contest_name"`
	StartTime   time.Time `json:"start_time"`
	// must be greater than start time
	EndTime time.Time `json:"end_time"`
	// must be equal to difference between end time and start time
	Duration          time.Duration `json:"duration"`
	RegistrationStart time.Time     `json:"registration_start"`
	// must be greater than registration_start
	RegistrationEnd time.Time `json:"registration_end"`
	// should be created automatically
	AnnouncementBlog int64 `json:"announcement_blog"`
	// should be created automatically
	EditorialBlog int64        `json:"editorial_blog"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
}

type ContestCreator struct {
	ContestID int64     `json:"contest_id"`
	CreatorID int64     `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ContestRegistered struct {
	ContestID int64     `json:"contest_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ProbTag struct {
	ProblemID int64  `json:"problem_id"`
	Tag       string `json:"tag"`
}

type Problem struct {
	ID           int64  `json:"id"`
	ProblemName  string `json:"problem_name"`
	Description  string `json:"description"`
	SampleInput  string `json:"sample_input"`
	SampleOutput string `json:"sample_output"`
	// to generate the output files in problemTests
	IdealSolution string `json:"ideal_solution"`
	// should be in seconds
	TimeLimit int32 `json:"time_limit"`
	// should be in MB
	MemoryLimit int32 `json:"memory_limit"`
	// should be in KB
	CodeSize  int32         `json:"code_size"`
	Rating    sql.NullInt32 `json:"rating"`
	CreatedAt time.Time     `json:"created_at"`
	ContestID int64         `json:"contest_id"`
}

type ProblemCreator struct {
	ProblemID int64     `json:"problem_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type ProblemTest struct {
	ID        int64  `json:"id"`
	ProblemID int64  `json:"problem_id"`
	Input     string `json:"input"`
	Output    string `json:"output"`
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
