-- name: GetUsers :many
SELECT * FROM Users;

-- name: CreateUser :one
INSERT INTO Users (name, dob)
VALUES ($1,$2)
RETURNING *;

-- name: GetUserByID :one
SELECT * from Users WHERE id = $1;

-- name: UpdateUser :one
UPDATE Users
SET 
   name = COALESCE(sqlc.narg('name'),name),
   dob = COALESCE(sqlc.narg('dob'),dob)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: ListAllUsers :many
SELECT * from Users;

-- name: GetUsersPaginated :many
SELECT * FROM Users
ORDER BY id
OFFSET $1 LIMIT $2;

-- name: CountUsers :one
SELECT COUNT(*) FROM Users;
