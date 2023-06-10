// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: blog.sql

package db

import (
	"context"
	"time"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blogs (
  blog_title,
  blog_content,
  created_by,
  publish_at
) VALUES (
  $1, $2, $3, $4
) RETURNING id, blog_title, blog_content, created_by, created_at, publish_at, votes_count
`

type CreateBlogParams struct {
	BlogTitle   string    `json:"blog_title"`
	BlogContent string    `json:"blog_content"`
	CreatedBy   string    `json:"created_by"`
	PublishAt   time.Time `json:"publish_at"`
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog,
		arg.BlogTitle,
		arg.BlogContent,
		arg.CreatedBy,
		arg.PublishAt,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.BlogTitle,
		&i.BlogContent,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.PublishAt,
		&i.VotesCount,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT id, blog_title, blog_content, created_by, created_at, publish_at, votes_count FROM blogs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBlog(ctx context.Context, id int64) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.BlogTitle,
		&i.BlogContent,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.PublishAt,
		&i.VotesCount,
	)
	return i, err
}

const listBlogs = `-- name: ListBlogs :many
SELECT id, blog_title, blog_content, created_by, created_at, publish_at, votes_count FROM blogs
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListBlogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBlogs(ctx context.Context, arg ListBlogsParams) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, listBlogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Blog{}
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.BlogTitle,
			&i.BlogContent,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.PublishAt,
			&i.VotesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBlog = `-- name: UpdateBlog :one
UPDATE blogs
  set blog_title = $2,
  blog_content = $3,
  publish_at = $4
WHERE id = $1
RETURNING id, blog_title, blog_content, created_by, created_at, publish_at, votes_count
`

type UpdateBlogParams struct {
	ID          int64     `json:"id"`
	BlogTitle   string    `json:"blog_title"`
	BlogContent string    `json:"blog_content"`
	PublishAt   time.Time `json:"publish_at"`
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateBlog,
		arg.ID,
		arg.BlogTitle,
		arg.BlogContent,
		arg.PublishAt,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.BlogTitle,
		&i.BlogContent,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.PublishAt,
		&i.VotesCount,
	)
	return i, err
}

const updateblog = `-- name: Updateblog :one
UPDATE blogs
  set blog_title = $2,
  blog_content = $3
WHERE id = $1
RETURNING id, blog_title, blog_content, created_by, created_at, publish_at, votes_count
`

type UpdateblogParams struct {
	ID          int64  `json:"id"`
	BlogTitle   string `json:"blog_title"`
	BlogContent string `json:"blog_content"`
}

func (q *Queries) Updateblog(ctx context.Context, arg UpdateblogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateblog, arg.ID, arg.BlogTitle, arg.BlogContent)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.BlogTitle,
		&i.BlogContent,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.PublishAt,
		&i.VotesCount,
	)
	return i, err
}