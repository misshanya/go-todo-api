-- name: CreateTodo :exec
INSERT INTO todo (
    title, content
) VALUES (
    $1, $2
);

-- name: GetTodoByID :one
SELECT * FROM todo
WHERE id = $1 LIMIT 1;

-- name: ListTodosByCreatedAt :many
SELECT * FROM todo
ORDER BY created_at DESC;

-- name: ListTodosByUpdatedAt :many
SELECT * FROM todo
ORDER BY updated_at DESC;

-- name: UpdateTodo :exec
UPDATE todo
SET 
    title = COALESCE($2, title),
    content = COALESCE($3, content),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteTodo :exec
DELETE FROM todo
WHERE id = $1;
