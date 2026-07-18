-- name: CreateInstructor :one
INSERT INTO instructors (id, created_at, updated_at, email, instructor_name, password_hash)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: RemoveInstructors :exec
DELETE FROM instructors;
