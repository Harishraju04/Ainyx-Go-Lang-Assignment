-- name: GetUsers :many
SELECT * FROM Users;

-- name: CreateUser :one
INSERT INTO Users (name, dob)
VALUES ($1,$2)
RETURNING *;