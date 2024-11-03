package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/misshanya/go-todo-api/internal/db"
)

type TodoService struct {
	queries *db.Queries
}

func NewTodoService(queries *db.Queries) *TodoService {
	return &TodoService{queries: queries}
}

// CreateTodo creates new todo task
func (s *TodoService) CreateTodo(ctx context.Context, title, content string) error {
	params := db.CreateTodoParams{
		Title:   pgtype.Text{String: title, Valid: true},
		Content: pgtype.Text{String: content, Valid: true},
	}
	return s.queries.CreateTodo(ctx, params)
}

// GetTodoByID returns todo task by its ID
func (s *TodoService) GetTodoByID(ctx context.Context, id int32) (db.Todo, error) {
	return s.queries.GetTodoByID(ctx, id)
}

// UpdateTodo updates todo task by its ID
func (s *TodoService) UpdateTodo(ctx context.Context, id int32, title, content *string) error {
	params := db.UpdateTodoParams{
		ID:      id,
		Title:   pgtype.Text{String: "", Valid: false},
		Content: pgtype.Text{String: "", Valid: false},
	}
	if title != nil {
		params.Title = pgtype.Text{String: *title, Valid: true}
	}
	if content != nil {
		params.Content = pgtype.Text{String: *content, Valid: true}
	}
	return s.queries.UpdateTodo(ctx, params)
}

// DeleteTodo deletes todo task by its ID
func (s *TodoService) DeleteTodo(ctx context.Context, id int32) error {
	return s.queries.DeleteTodo(ctx, id)
}

// ListTodosByCreatedAt returns a slice of todos
// sorted by create time in descending order
func (s *TodoService) ListTodosByCreatedAt(ctx context.Context) ([]db.Todo, error) {
	return s.queries.ListTodosByCreatedAt(ctx)
}

// ListTodosByUpdatedAt returns a slice of todos
// sorted by update time in descending order
func (s *TodoService) ListTodosByUpdatedAt(ctx context.Context) ([]db.Todo, error) {
	return s.queries.ListTodosByUpdatedAt(ctx)
}
