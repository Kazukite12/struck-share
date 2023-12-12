-- name: CreateFollower :one
INSERT INTO follower (
 user_id, 
 followed_user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListFollower :many
SELECT * FROM follower
WHERE user_id = $1
LIMIT $2
OFFSET $3;

-- name: GetUserFollower :one
SELECT * FROM follower
WHERE user_id = $1 AND followed_user_id = $2
LIMIT 1;

-- name: DeleteFollower :exec
DELETE FROM follower
WHERE followed_user_id = $1;

-- name: CountUserTotalFollower :many
SELECT count(*) FROM follower
WHERE user_id = $1;


