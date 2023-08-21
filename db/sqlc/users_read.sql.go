// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users_read.sql

package db

import (
	"context"
	"database/sql"
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

const getFollowersByUserID = `-- name: GetFollowersByUserID :many
SELECT users.id, username, users.created_at, users.deleted_at, followers.id, user_id, follower_id, followers.created_at, followers.deleted_at from users
JOIN followers on followers.user_id = users.id
WHERE users.id = $1
`

type GetFollowersByUserIDRow struct {
	ID          uuid.UUID
	Username    string
	CreatedAt   time.Time
	DeletedAt   sql.NullTime
	ID_2        uuid.UUID
	UserID      uuid.UUID
	FollowerID  uuid.UUID
	CreatedAt_2 time.Time
	DeletedAt_2 sql.NullTime
}

func (q *Queries) GetFollowersByUserID(ctx context.Context, id uuid.UUID) ([]GetFollowersByUserIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowersByUserID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFollowersByUserIDRow{}
	for rows.Next() {
		var i GetFollowersByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.CreatedAt,
			&i.DeletedAt,
			&i.ID_2,
			&i.UserID,
			&i.FollowerID,
			&i.CreatedAt_2,
			&i.DeletedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsers = `-- name: GetUsers :many
SELECT id, username, created_at, deleted_at from users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.CreatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
