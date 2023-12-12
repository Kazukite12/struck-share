// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: post_media.sql

package db

import (
	"context"
)

const createMedia = `-- name: CreateMedia :one
INSERT INTO post_media (
  post_id,
  media_file
) VALUES (
  $1, $2
)
RETURNING media_id, post_id, media_file, created_at
`

type CreateMediaParams struct {
	PostID    int64
	MediaFile string
}

func (q *Queries) CreateMedia(ctx context.Context, arg CreateMediaParams) (PostMedium, error) {
	row := q.db.QueryRowContext(ctx, createMedia, arg.PostID, arg.MediaFile)
	var i PostMedium
	err := row.Scan(
		&i.MediaID,
		&i.PostID,
		&i.MediaFile,
		&i.CreatedAt,
	)
	return i, err
}

const getMedia = `-- name: GetMedia :one
SELECT media_id, post_id, media_file, created_at FROM post_media
WHERE post_id = $1 LIMIT 1
`

func (q *Queries) GetMedia(ctx context.Context, postID int64) (PostMedium, error) {
	row := q.db.QueryRowContext(ctx, getMedia, postID)
	var i PostMedium
	err := row.Scan(
		&i.MediaID,
		&i.PostID,
		&i.MediaFile,
		&i.CreatedAt,
	)
	return i, err
}
