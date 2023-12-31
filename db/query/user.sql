-- name: CreateUser :one
INSERT INTO users (
  username,
  profile_picture,
  full_name,
  hashed_password,
  email
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

