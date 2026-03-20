-- name: CreateList :one
INSERT INTO lists (
	folder, branch
) VALUES (
	?, ?
) RETURNING *;

-- name: CreateTodo :one
INSERT INTO todos (
	list_id, description
) VALUES (
	?, ?
) RETURNING *;

-- name: GetTodosByListId :many
SELECT * FROM todos
WHERE list_id = ?;

-- name: GetTodoStatusById :one
SELECT done FROM todos
WHERE id = ?;

-- name: GetListIdByFolderBranch :one
SELECT id FROM lists 
WHERE folder = ? AND branch = ? LIMIT 1;

-- name: GetListById :one
SELECT * FROM lists
WHERE id = ? LIMIT 1;

-- name: UpdateTodoStatus :exec
UPDATE todos
SET done = ?
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

-- name: DeleteList :exec
DELETE FROM lists
WHERE id = ?;

