-- name: CreateUser :one
INSERT INTO users (id, email, username, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;