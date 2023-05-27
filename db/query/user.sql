-- name: CreateUser :one
INSERT INTO users (
  name,
  username,
  email,
  password,
  dob
) VALUES (
  $1, $2, $3, $4, $5
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
  set name = $2,
  email = $3,
  password = $4,
  profileimg = $5,
  motto = $6,
  dob = $7,
  is_setter = $8
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

