-- name: CreateBlog :one
INSERT INTO blogs (
  blog_title,
  blog_content,
  created_by,
  publish_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;


-- name: GetBlog :one
SELECT * FROM blogs
WHERE id = $1 LIMIT 1;

-- name: ListBlogs :many
SELECT * FROM blogs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;

-- name: UpdateBlog :one
UPDATE blogs
SET
  blog_title = COALESCE(sqlc.narg(blog_title), blog_title),
  blog_content = COALESCE(sqlc.narg(blog_content), blog_content),
  publish_at = COALESCE(sqlc.narg(publish_at), publish_at)
WHERE id = sqlc.arg(id)
RETURNING *;
