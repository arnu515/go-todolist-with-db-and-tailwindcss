-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(id, username, password) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;



-- name: GetLists :many
SELECT * FROM lists WHERE user_id = $1;

-- name: GetList :one
SELECT * FROM lists WHERE id = $1 AND user_id = $2 LIMIT 1;

-- name: CreateList :one
INSERT INTO lists(id, title, user_id) VALUES ($1, $2, $3) RETURNING *; 

-- name: UpdateList :exec
UPDATE lists SET title = $2 WHERE id = $1 AND user_id = $3;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1 AND user_id = $2;



-- name: GetTodos :many
SELECT * FROM todos WHERE user_id = $1 AND list_id = $2;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1 AND list_id = $2 AND user_id = $3 LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos(id, text, list_id, user_id) VALUES ($1, $2, $3, $4) RETURNING *; 

-- name: UpdateTodo :exec
UPDATE todos SET text = $2, completed = $3 WHERE id = $1 AND list_id = $4 AND user_id = $5;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1 AND list_id = $2 AND user_id = $3;
