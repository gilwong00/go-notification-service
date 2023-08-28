-- name: ListSendableNotifications :many
SELECT
	notification_queue.message,
	notification_queue.attempts,
	users.url
FROM notification_queue
JOIN followers ON folowers.id = notification_queue.follower_id
JOIN users ON users.id = folowers.follower_id
JOIN notification_state ON notification_state.id = notification_queue.state_id
WHERE users.url IS NOT NULL AND notification_state.state != 'success'
ORDER BY notification_queue.created_at DESC;
