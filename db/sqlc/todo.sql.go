// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: todo.sql

package database

import (
	"context"
	"time"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (name, description, username, deadline)
VALUES ($1, $2, $3, $4)
RETURNING id, name, description, username, created_at, deadline
`

type CreateTodoParams struct {
	Name        string
	Description string
	Username    string
	Deadline    time.Time
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.Name,
		arg.Description,
		arg.Username,
		arg.Deadline,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Username,
		&i.CreatedAt,
		&i.Deadline,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE
FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, name, description, username, created_at, deadline
FROM todos
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Username,
		&i.CreatedAt,
		&i.Deadline,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, name, description, username, created_at, deadline
FROM todos
WHERE username = $1
ORDER BY deadline
`

func (q *Queries) ListTodos(ctx context.Context, username string) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Username,
			&i.CreatedAt,
			&i.Deadline,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
