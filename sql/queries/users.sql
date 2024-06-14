-- name: CreateUser :one
INSERT INTO users (id, user_name, email, password, created_at, updated_at, currency)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;