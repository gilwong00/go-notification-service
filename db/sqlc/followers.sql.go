// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: followers.sql

package db

import (
	"context"

	uuid "github.com/gofrs/uuid/v5"
)

const createFollower = `-- name: CreateFollower :one
INSERT INTO followers (
	user_id,
	follower_id
) VALUES (
	$1,
	$2
) RETURNING id, user_id, follower_id, created_at, deleted_at
`

type CreateFollowerParams struct {
	UserID     uuid.UUID
	FollowerID uuid.UUID
}

func (q *Queries) CreateFollower(ctx context.Context, arg CreateFollowerParams) (Follower, error) {
	row := q.db.QueryRowContext(ctx, createFollower, arg.UserID, arg.FollowerID)
	var i Follower
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FollowerID,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
