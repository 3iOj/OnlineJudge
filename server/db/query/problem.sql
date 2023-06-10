-- name: CreateProblem :one
INSERT INTO problems (
  problem_name,
  description,
  sample_input,
  sample_output,
  ideal_solution,
  time_limit,
  memory_limit,
  code_size,
  created_at,
  contest_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetProblem :one
SELECT * FROM Problems
WHERE id = $1 LIMIT 1;

-- name: ListProblems :many
SELECT * FROM Problems
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProblem :one
UPDATE Problems
  set problem_name = $2,
  description = $3,
  sample_input = $4,
  sample_output = $5,
  ideal_solution = $6,
  time_limit = $7,
  memory_limit = $8,
  code_size = $9,
  rating = $10
WHERE id = $1
RETURNING *;

-- name: DeleteProblem :exec
DELETE FROM Problems
WHERE id = $1;


/*
TO-DO
adding problem tags
searching problems - tags, name, rating, ... ?
*/