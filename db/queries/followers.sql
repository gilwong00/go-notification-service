-- name: CreateFollower :one
INSERT INTO followers (
	user_id,
	follower_id
) VALUES (
	$1,
	$2
) RETURNING *;

