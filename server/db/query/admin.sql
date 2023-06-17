-- name: CreateAdmin :one
INSERT INTO admin (
  name,
  username,
  email,
  password
) VALUES (
  $1, $2, $3, $4
) RETURNING *;


-- name: GetAdmin :one
SELECT * FROM admin
WHERE username = $1 LIMIT 1;

-- name: ListAdmins :many
SELECT * FROM admin
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAdmin :one
UPDATE admin
SET 
  password = COALESCE(sqlc.narg(password), password),
  name = COALESCE(sqlc.narg(name), name),
  email = COALESCE(sqlc.narg(email), email)
WHERE
  username = sqlc.arg(username)
RETURNING *;

-- name: DeleteAdmin :exec
DELETE FROM admin
WHERE username = $1;

