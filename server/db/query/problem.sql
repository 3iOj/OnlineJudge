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
SET 
  problem_name = COALESCE(sqlc.narg(problem_name), problem_name),
  description = COALESCE(sqlc.narg(description), description),
  sample_input = COALESCE(sqlc.narg(sample_input), sample_input),
  sample_output = COALESCE(sqlc.narg(sample_output), sample_output),
  ideal_solution = COALESCE(sqlc.narg(ideal_solution), ideal_solution),
  time_limit = COALESCE(sqlc.narg(time_limit), time_limit),
  memory_limit = COALESCE(sqlc.narg(memory_limit), memory_limit),
  code_size = COALESCE(sqlc.narg(code_size), code_size),
  rating = COALESCE(sqlc.narg(rating), rating),
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