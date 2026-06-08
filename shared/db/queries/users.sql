-- name: CreateUser :one
INSERT INTO users(user_name, full_name)
VALUES($1, $2)
RETURNING id;

-- name: DeleteUserById :execrows
DELETE FROM users
WHERE id = $1;

-- name: GetUserById :one
SELECT user_name, full_name, created_at
FROM users
WHERE id = $1;
