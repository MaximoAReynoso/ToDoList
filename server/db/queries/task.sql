-- name: CreateTask :one
INSERT INTO task (title, description, completed) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTask :one
SELECT * FROM task WHERE id = $1;

-- name: ListTasks :many
SELECT * FROM task;

-- name: UpdateTask :exec
UPDATE task SET title = $2, description = $3, completed = $4 WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM task WHERE id = $1;
