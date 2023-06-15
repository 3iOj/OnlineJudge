-- name: CreateContest :one
INSERT INTO contests (
  contest_name,
  duration
) VALUES (
  $1, $2
) RETURNING *;


-- name: GetContest :one
SELECT * FROM contests as C INNER JOIN problems as P ON
C.id = P.contest_id WHERE C.id = $1 ;

-- name: AddContestCreators :one
INSERT INTO contest_creators (
    contest_id,
    creator_name
) VALUES (
    $1, $2
) RETURNING *;

-- name: DeleteContestCreators :exec
DELETE FROM contest_creators
WHERE contest_id = $1;

-- name: ListContests :many
SELECT * FROM contests
WHERE ispublish IS TRUE
LIMIT $1
OFFSET $2;

-- name: DeleteContest :exec
DELETE FROM contests
WHERE id = $1;

-- name: UpdateContest :one
UPDATE contests
SET
  contest_name = COALESCE(sqlc.narg(contest_name), contest_name),
  start_time = COALESCE(sqlc.narg(start_time), start_time),
  end_time = COALESCE(sqlc.narg(end_time), end_time),
  duration = COALESCE(sqlc.narg(duration), duration),
  registration_start = COALESCE(sqlc.narg(registration_start), registration_start),
  registration_end = COALESCE(sqlc.narg(registration_end), end_time),
  announcement_blog = COALESCE(sqlc.narg(announcement_blog), announcement_blog),
  editorial_blog = COALESCE(sqlc.narg(editorial_blog), editorial_blog),
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
  ispublish = COALESCE(sqlc.narg(ispublish), ispublish)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: AddParticipant :one
INSERT INTO contest_registered (
  contest_id,
  username
) VALUES (
  $1, $2
) RETURNING *;


-- name: DeleteParticipant :exec
DELETE FROM contest_registered
WHERE username = $1;