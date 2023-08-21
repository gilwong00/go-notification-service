-- name: GetFollowersByUserID :many
SELECT * from users
JOIN followers on followers.user_id = users.id
WHERE users.id = $1;
