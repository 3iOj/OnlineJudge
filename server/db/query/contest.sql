-- name: CreateContest :one
INSERT INTO contests (
  contest_name,
  start_time,
  end_time,
  duration,
  registration_start,
  registration_end,
  announcement_blog,
  editorial_blog
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;


-- name: GetContest :one
SELECT * FROM contests
WHERE id = $1 LIMIT 1;

-- name: AddContestCreators :one
INSERT INTO contest_creators (
    contest_id,
    creator_id
) VALUES (
    $1, $2
) RETURNING *;


-- name: ListContests :many
SELECT * FROM contests
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
-- UPDATE users
--   set name = $2,
--   email = $3,
--   password = $4,
--   profileimg = $5,
--   motto = $6,
--   dob = $7,
--   is_setter = $8
-- WHERE username = $1
-- RETURNING *;

-- -- name: DeleteUser :exec
-- DELETE FROM users
-- WHERE username = $1;
