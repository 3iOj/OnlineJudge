package db

import (
	"context"
	"database/sql"
	"time"
)

type CreateContestTxParams struct {
	ContestName       string    `json:"contest_name"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	Duration          int64     `json:"duration"`
	RegistrationStart time.Time `json:"registration_start"`
	RegistrationEnd   time.Time `json:"registration_end"`
	AnnouncementBlog  int64     `json:"announcement_blog"`
	EditorialBlog     int64     `json:"editorial_blog"`
}

type CreateContestTxResponse struct {
	Contest          Contest `json:"contest"`
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateContestTxParams) (CreateContestTxResponse, error) {
	var rsp CreateContestTxResponse

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		announcementBlog, err := q.UpdateBlog(ctx,UpdateBlogParams{
				PublishAt: sql.NullTime{Time: time.Now(), Valid: true},
				ID:arg.AnnouncementBlog,
		})
		if err != nil {
			return err
		}

		editorialBlog, err := q.UpdateBlog(ctx,
			UpdateBlogParams{
				PublishAt: sql.NullTime{Time: time.Now(), Valid: true},
				ID:arg.EditorialBlog,
		})
		if err != nil {
			return err
		}

		rsp.Contest, err = q.CreateContest(ctx,
			CreateContestParams{
				ContestName:       arg.ContestName,
				StartTime:         arg.StartTime,
				EndTime:           arg.EndTime,
				Duration:          arg.Duration,
				RegistrationStart: arg.RegistrationStart,
				RegistrationEnd:   arg.RegistrationEnd,
				AnnouncementBlog:  announcementBlog.ID,
				EditorialBlog:     editorialBlog.ID,
			})
		if err != nil {
			return err
		}
		return err
	})
	return rsp, err
}
