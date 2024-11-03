// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTodo = `-- name: CreateTodo :exec
INSERT INTO todo (
    title, content
) VALUES (
    $1, $2
)
`

type CreateTodoParams struct {
	Title   pgtype.Text
	Content pgtype.Text
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) error {
	_, err := q.db.Exec(ctx, createTodo, arg.Title, arg.Content)
	return err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todo
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteTodo, id)
	return err
}

const getTodoByID = `-- name: GetTodoByID :one
SELECT id, title, content, done, created_at, updated_at FROM todo
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodoByID(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodoByID, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Done,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTodosByCreatedAt = `-- name: ListTodosByCreatedAt :many
SELECT id, title, content, done, created_at, updated_at FROM todo
ORDER BY created_at DESC
`

func (q *Queries) ListTodosByCreatedAt(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodosByCreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Done,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTodosByUpdatedAt = `-- name: ListTodosByUpdatedAt :many
SELECT id, title, content, done, created_at, updated_at FROM todo
ORDER BY updated_at DESC
`

func (q *Queries) ListTodosByUpdatedAt(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodosByUpdatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Done,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :exec
UPDATE todo
SET 
    title = COALESCE($2, title),
    content = COALESCE($3, content),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateTodoParams struct {
	ID      int32
	Title   pgtype.Text
	Content pgtype.Text
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.Exec(ctx, updateTodo, arg.ID, arg.Title, arg.Content)
	return err
}