-- name: CreateNotificationState :one
INSERT INTO notification_state (
	state
) VALUES (
	$1
) RETURNING *;

-- name: UpdateNotificationStateByID :exec
UPDATE notification_state
SET state = $2, message = $3
WHERE id = $1;

-- name: CreateNotificationEvent :one
INSERT INTO notification_queue (
	message,
	follower_id,
	state_id,
	attempts
) VALUES (
	$1, $2, $3, $4
) RETURNING *;
