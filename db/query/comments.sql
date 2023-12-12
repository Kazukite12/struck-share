-- name: CreateComment :one
INSERT INTO comments (
  created_by_user_id,
  post_id,
  comment
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE comment_id = $1 LIMIT 1;

-- name: ListComment :many
SELECT * FROM comments
WHERE post_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE comment_id = $1;