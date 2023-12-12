-- name: CreateMedia :one
INSERT INTO post_media (
  post_id,
  media_file
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetMedia :one
SELECT * FROM post_media
WHERE post_id = $1 LIMIT 1;