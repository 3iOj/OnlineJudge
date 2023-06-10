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

-- name: Updateblog :one
UPDATE blogs
  set blog_title = $2,
  blog_content = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;

-- name: UpdateBlog :one
UPDATE blogs
  set blog_title = $2,
  blog_content = $3,
  publish_at = $4
WHERE id = $1
RETURNING *;
