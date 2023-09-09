-- name: CreateSubmission :one
INSERT INTO Submissions (
  problem_id,
  username,
  user_id,
  contest_id,
  language,
  code,
  submitted_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSubmission :one
SELECT  FROM Submissions
WHERE id = $1 LIMIT 1;

-- name: ListSubmissions :many
SELECT * FROM Submissions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSubmission :one
UPDATE Submissions
SET 
  verdict = COALESCE(sqlc.narg(verdict), verdict),
  memory_consumed = COALESCE(sqlc.narg(memory_consumed), memory_consumed),
  exec_time = COALESCE(sqlc.narg(exec_time), exec_time),
  score = COALESCE(sqlc.narg(score), score)
WHERE id = $1
RETURNING *;

-- name: DeleteSubmission :exec
DELETE FROM Submissions
WHERE id = $1;
