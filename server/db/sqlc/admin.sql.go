// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: admin.sql

package db

import (
	"context"
	"database/sql"
)

const createAdmin = `-- name: CreateAdmin :one
INSERT INTO admin (
  name,
  username,
  email,
  password
) VALUES (
  $1, $2, $3, $4
) RETURNING id, name, username, email, password, created_at
`

type CreateAdminParams struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateAdmin(ctx context.Context, arg CreateAdminParams) (Admin, error) {
	row := q.db.QueryRowContext(ctx, createAdmin,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAdmin = `-- name: DeleteAdmin :exec
DELETE FROM admin
WHERE username = $1
`

func (q *Queries) DeleteAdmin(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteAdmin, username)
	return err
}

const getAdmin = `-- name: GetAdmin :one
SELECT id, name, username, email, password, created_at FROM admin
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetAdmin(ctx context.Context, username string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdmin, username)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const listAdmins = `-- name: ListAdmins :many
SELECT id, name, username, email, password, created_at FROM admin
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAdminsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAdmins(ctx context.Context, arg ListAdminsParams) ([]Admin, error) {
	rows, err := q.db.QueryContext(ctx, listAdmins, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Admin{}
	for rows.Next() {
		var i Admin
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdmin = `-- name: UpdateAdmin :one
UPDATE admin
SET 
  password = COALESCE($1, password),
  name = COALESCE($2, name),
  email = COALESCE($3, email)
WHERE
  username = $4
RETURNING id, name, username, email, password, created_at
`

type UpdateAdminParams struct {
	Password sql.NullString `json:"password"`
	Name     sql.NullString `json:"name"`
	Email    sql.NullString `json:"email"`
	Username string         `json:"username"`
}

func (q *Queries) UpdateAdmin(ctx context.Context, arg UpdateAdminParams) (Admin, error) {
	row := q.db.QueryRowContext(ctx, updateAdmin,
		arg.Password,
		arg.Name,
		arg.Email,
		arg.Username,
	)
	var i Admin
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
