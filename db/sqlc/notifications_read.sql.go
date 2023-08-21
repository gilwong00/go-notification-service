// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: notifications_read.sql

package db

import (
	"context"
)

const listNotifications = `-- name: ListNotifications :many
SELECT id, message, follower_id, state_id, attempts, created_at, updated_at FROM notification_queue
`

func (q *Queries) ListNotifications(ctx context.Context) ([]NotificationQueue, error) {
	rows, err := q.db.QueryContext(ctx, listNotifications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []NotificationQueue{}
	for rows.Next() {
		var i NotificationQueue
		if err := rows.Scan(
			&i.ID,
			&i.Message,
			&i.FollowerID,
			&i.StateID,
			&i.Attempts,
			&i.CreatedAt,
			&i.UpdatedAt,
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
