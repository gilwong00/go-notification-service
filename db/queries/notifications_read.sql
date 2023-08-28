-- name: ListSendableNotifications :many
SELECT
	notification_queue.id,
	notification_queue.message,
	notification_queue.attempts,
	notification_state.id AS notification_state_id,
	users.url
FROM notification_queue
JOIN followers ON followers.user_id = notification_queue.follower_id
JOIN users ON users.id = followers.follower_id
JOIN notification_state ON notification_state.id = notification_queue.state_id
WHERE
	users.url IS NOT NULL AND
	notification_state.state != 'success' AND
	notification_state.completed_at IS NULL
ORDER BY notification_queue.created_at DESC;
