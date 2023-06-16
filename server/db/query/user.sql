-- name: CreateUser :one
INSERT INTO users (
  name,
  username,
  email,
  password,
  dob,
  profileimg,
  motto
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET 
  password = COALESCE(sqlc.narg(password), password),
  name = COALESCE(sqlc.narg(name), name),
  email = COALESCE(sqlc.narg(email), email),
  dob = COALESCE(sqlc.narg(dob), dob),
  profileimg = COALESCE(sqlc.narg(profileimg), profileimg),
  motto = COALESCE(sqlc.narg(motto), motto),
  is_setter = COALESCE(sqlc.narg(is_setter), is_setter)
WHERE
  username = sqlc.arg(username)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

