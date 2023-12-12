-- name: CreatePost :one
INSERT INTO posts (
  created_by_user_id,
  caption
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE post_id = $1 LIMIT 1;

-- name: ListPost :many
SELECT * FROM posts
WHERE created_by_user_id = $1
ORDER BY post_id
LIMIT $2
OFFSET $3;

-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = $1;

