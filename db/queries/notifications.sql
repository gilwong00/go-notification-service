-- name: CreateNotificationState :one
INSERT INTO notification_state (
	state
) VALUES (
	$1
) RETURNING *;

-- name: UpdateNotificationStateByID :exec
UPDATE notification_state
SET
	state = $2,
	message = $3,
	completed_at = $4
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

-- name: UpdateNotificationAttemptCount :exec
UPDATE notification_queue
SET attempts = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteNotificationEvent :exec
UPDATE notification_state
SET
	state = $2,
	message = $3,
	completed_at = CURRENT_TIMESTAMP
WHERE id = $1;
