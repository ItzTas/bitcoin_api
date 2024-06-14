// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, user_name, email, password, created_at, updated_at, currency)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_name, email, password, created_at, updated_at, currency
`

type CreateUserParams struct {
	ID        uuid.UUID
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Currency  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.UserName,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Currency,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Currency,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, user_name, email, password, created_at, updated_at, currency FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Currency,
	)
	return i, err
}
