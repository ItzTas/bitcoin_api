-- name: CreateUser :one
INSERT INTO users (id, user_name, email, password, created_at, updated_at, currency)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users
LIMIT $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET password = $1, email = $2, user_name = $3
WHERE id = $4
RETURNING *;