-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at, email, password_hash)
VALUES (
    gen_random_uuid(),
    $1,
    NOW(),
    NOW(),
    $2,
    $3
)
RETURNING *;

-- name: RemoveUsers :exec
DELETE FROM users;

-- name: QueryUserEmail :one
SELECT * FROM users WHERE email = $1;