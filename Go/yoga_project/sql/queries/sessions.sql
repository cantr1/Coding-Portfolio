-- name: CreateSession :one
INSERT INTO sessions (id, created_at, updated_at, start_time, end_time, instructor_id, difficulty, class_size, description)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: RemoveSessions :exec
DELETE FROM sessions;

-- name: QuerySessionID :one
SELECT * FROM sessions WHERE id = $1;

-- name: QueryAvailableSessionsInstructor :many
SELECT * FROM sessions WHERE instructor_id = $1;

-- name: QueryAvailableSessionsDifficulty :many
SELECT * FROM sessions WHERE difficulty = $1;

-- name: QuerySessionsInstructor :many
SELECT * FROM sessions WHERE instructor_id = $1;

-- name: QuerySessionsDifficulty :many
SELECT * FROM sessions WHERE difficulty = $1;

-- name: GetAllSessions :many
SELECT * FROM sessions;