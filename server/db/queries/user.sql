-- name: CreateUser :one
INSERT INTO task (id, title, description) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM task WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM task;

-- name: UpdateUser :exec
UPDATE task SET title = $2, description = $3 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM task WHERE id = $1;
