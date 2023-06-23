// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id, full_name, email, gender, date_of_birth, created_at, country_code, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, full_name, email, gender, date_of_birth, created_at, updated_at, country_code
`

type CreateUserParams struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
	CountryCode string    `json:"country_code"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Gender,
		arg.DateOfBirth,
		arg.CreatedAt,
		arg.CountryCode,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Gender,
		&i.DateOfBirth,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CountryCode,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, full_name, email, gender, date_of_birth, created_at, updated_at, country_code
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Gender,
		&i.DateOfBirth,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CountryCode,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, full_name, email, gender, date_of_birth, created_at, updated_at, country_code
FROM users
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.Gender,
			&i.DateOfBirth,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CountryCode,
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

const updateUsers = `-- name: UpdateUsers :exec
UPDATE users
SET full_name = $2
WHERE id = $1
RETURNING id, full_name, email, gender, date_of_birth, created_at, updated_at, country_code
`

type UpdateUsersParams struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}

func (q *Queries) UpdateUsers(ctx context.Context, arg UpdateUsersParams) error {
	_, err := q.db.ExecContext(ctx, updateUsers, arg.ID, arg.FullName)
	return err
}
