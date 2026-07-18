-- +goose Up
CREATE TABLE instructors (
    id UUID PRIMARY KEY, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email TEXT NOT NULL UNIQUE,
    instructor_name TEXT NOT NULL,
    password_hash TEXT NOT NULL
);

-- +goose Down
DROP TABLE instructors;
