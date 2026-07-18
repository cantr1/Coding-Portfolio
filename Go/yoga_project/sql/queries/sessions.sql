-- name: CreateSession :one
INSERT INTO sessions (id, created_at, updated_at, start_time, end_time, instructor, difficulty)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: RemoveSessions :exec
DELETE FROM sessions;

-- name: QuerySessionsInstructor :many
SELECT * FROM sessions WHERE instructor = $1;

-- name: QuerySessionsDifficulty :many
SELECT * FROM sessions WHERE difficulty = $1;