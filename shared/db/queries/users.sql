-- name: CreateUser :one
INSERT INTO users(user_name, full_name)
VALUES($1, $2)
RETURNING id;
