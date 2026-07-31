/*
A design problem in SQL is what to do about data in a table
that references a key from another table when said reference
is deleted

Cascade is generally best used on data that is not
important to keep should the reference be deleted

This file is a good example, having users and a
user sessions table -> not worth keeping without the user
*/

-- Create primary table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    display_name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create secondary table
CREATE TABLE user_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    session_name TEXT NOT NULL,
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ended_at TIMESTAMPTZ
);

-- Insert data into primary table
INSERT INTO users (name, email, password_hash, display_name)
VALUES (
    'Kelly Cantrell',
    'kelz@example.com',
    'abcd12345',
    'kelz'
),
(
    'Grace Hopper',
    'grace@example.com',
    'efghi6789',
    'gr8c3'
),
(
    'Ada Lovelace',
    'ada@example.com',
    'jklmno6789',
    'ada'
);

-- Insert data into reference table
INSERT INTO user_sessions (user_id, session_name)
SELECT id, 'Learn Ruby'
FROM users
WHERE email = 'kelz@example.com';

INSERT INTO user_sessions (user_id, session_name)
SELECT id, 'Build a Compiler'
FROM users
WHERE email = 'grace@example.com';

INSERT INTO user_sessions (user_id, session_name)
SELECT id, 'Think in Computer'
FROM users
WHERE email = 'ada@example.com';

-- End a session
UPDATE user_sessions
SET ended_at = NOW()
WHERE session_name = 'Think in Computer';

-- Get active sessions
SELECT * FROM user_sessions
WHERE ended_at IS NULL;

-- Get inactive sessions
SELECT * FROM user_sessions
WHERE ended_at IS NOT NULL;

-- Join to view the data we want from both tables - active sessions with a K in the name
SELECT 
u.name, 
s.session_name,
s.ended_at
FROM user_sessions AS s
JOIN users AS u 
    ON s.user_id = u.id
WHERE s.ended_at IS NULL AND u.name ILIKE 'k%';

-- Delete a user
DELETE FROM users WHERE name = 'Grace Hopper';

-- Check the reference table
SELECT * FROM user_sessions;

-- Cleanup
DROP TABLE user_sessions;
DROP TABLE users;
-- DROP DATABASE customer_db;