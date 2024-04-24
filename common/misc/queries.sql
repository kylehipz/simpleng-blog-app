-- name: CreateUser :one
INSERT INTO users (user_name) VALUES ($1) RETURNING *;

-- name: CreateFollowRel :one
INSERT INTO follow (follower, followee) VALUES ($1, $2) RETURNING *;
