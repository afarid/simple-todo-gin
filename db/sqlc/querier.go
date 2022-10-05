// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package database

import (
	"context"
)

type Querier interface {
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
	DeleteTodo(ctx context.Context, id int32) error
	GetTodo(ctx context.Context, id int32) (Todo, error)
	ListTodos(ctx context.Context, username string) ([]Todo, error)
}

var _ Querier = (*Queries)(nil)
