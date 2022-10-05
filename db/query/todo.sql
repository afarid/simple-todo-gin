-- name: GetTodo :one
SELECT *
FROM todos
WHERE id = $1
LIMIT 1;

-- name: ListTodos :many
SELECT *
FROM todos
WHERE username = $1
ORDER BY deadline;

-- name: CreateTodo :one
INSERT INTO todos (name, description, username, deadline)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteTodo :exec
DELETE
FROM todos
WHERE id = $1;