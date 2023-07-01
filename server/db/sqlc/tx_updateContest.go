package db

// import (
// 	"context"
// 	"database/sql"
// 	"time"
// )

// type UpdateContestTxParams struct {
// 	ID                int64     `json:"id" binding:"required"`
// 	ContestName       string    `json:"contest_name" binding:"required"`
// 	StartTime         time.Time `json:"start_time" binding:"required"`
// 	EndTime           time.Time `json:"end_time" binding:"required"`
// 	Duration          int64     `json:"duration" binding:"required"`
// 	RegistrationStart time.Time `json:"registration_start" binding:"required"`
// 	RegistrationEnd   time.Time `json:"registration_end" binding:"required"`
// 	AnnouncementBlog  int64     `json:"announcement_blog" binding:"required"`
// 	EditorialBlog     int64     `json:"editorial_blog" binding:"required"`
// 	ContestCreators   []string  `json:"contest_creators" binding:"required"`
// 	Ispublish         bool      `json:"ispublish"`
// }

// type UpdateContestTxResponse struct {
// 	Contest         Contest  `json:"contest"`
// 	ContestCreators []string `json:"contest_creators"`
// }

// func (store *SQLStore) UpdateContestTx(ctx context.Context, arg UpdateContestTxParams) (UpdateContestTxResponse, error) {
// 	var rsp UpdateContestTxResponse

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error
// 		var announcement_blog, editorial_blog Blog
// 		if arg.AnnouncementBlog != 0 {
// 			announcement_blog, err = q.UpdateBlog(ctx, UpdateBlogParams{
// 				PublishAt: sql.NullTime{Time: time.Now(), Valid: true},
// 				ID:        arg.AnnouncementBlog,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		if arg.EditorialBlog != 0 {
// 			editorial_blog, err = q.UpdateBlog(ctx,
// 				UpdateBlogParams{
// 					PublishAt: sql.NullTime{Time: time.Now(), Valid: true},
// 					ID:        arg.EditorialBlog,
// 				})
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		if arg.Ispublish == true {
// 			//checking all arguments should be NOT NULL
// 		}
// 		rsp.Contest, err = q.UpdateContest(ctx,
// 			UpdateContestParams{
// 				ContestName:       sql.NullString{String: arg.ContestName, Valid: true},
// 				EndTime:           sql.NullTime{Time: arg.EndTime, Valid: true},
// 				StartTime:         sql.NullTime{Time: arg.StartTime, Valid: true},
// 				Duration:          sql.NullInt64{Int64: arg.Duration, Valid: true},
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
// 				AddContestCreatorsParams{
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
