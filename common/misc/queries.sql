-- name: CreateUser :one
INSERT INTO users (user_name) VALUES ($1) RETURNING *;

-- name: CreateFollowRel :one
INSERT INTO follow (follower, followee) VALUES ($1, $2) RETURNING *;

-- name: DeleteFollowRel :exec
DELETE FROM follow where follower = $1 and followee = $2;

-- name: CreateBlog :one
INSERT INTO blogs (author, title, body) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateBlog :one
UPDATE blogs 
SET
  title = $2,
  body = $3
WHERE id = $1 RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs WHERE id = $1;

-- name: GetBlogsFromFollowees :many
SELECT b.* FROM blogs b 
INNER JOIN follow f ON b.author = f.followee
WHERE f.follower = $1 ORDER BY b.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetBlogsFromFolloweesCount :one
SELECT count(b.*) FROM blogs b 
INNER JOIN follow f ON b.author = f.followee
WHERE f.follower = $1;
