-- +goose Up
CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    instructor TEXT NOT NULL,
    difficulty INT NOT NULL
);

-- +goose Down
DROP TABLE sessions;