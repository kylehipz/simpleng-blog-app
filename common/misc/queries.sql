-- name: CreateUser :one
INSERT INTO users (user_name) VALUES ($1) RETURNING *;