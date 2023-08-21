-- name: CreateUser :one
INSERT INTO users (
	username
) VALUES (
	$1
) RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE users.id = $1;