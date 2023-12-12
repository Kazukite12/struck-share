-- name: CreatePostLikes :one
INSERT INTO post_likes (
  post_id,
  post_liker
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreatePostUnlikes :exec
DELETE FROM post_likes
WHERE post_liker = $1;

-- name: GetPostLike :one
SELECT * FROM post_likes
WHERE post_liker = $1 AND post_id = $2
LIMIT 1;

-- name: ListPostLiker :many
SELECT * FROM post_likes
WHERE post_id = $1
ORDER BY post_liker
LIMIT $2
OFFSET $3;

-- name: CountPostTotalLikes :many
SELECT count(*) FROM post_likes
WHERE post_id = $1;

