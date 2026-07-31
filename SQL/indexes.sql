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
),
(
    'John Wick',
    'john@example.com',
    'efghi6789',
    'mr.wick'
),
(
    'Vis Telimus',
    'vis@example.com',
    'jklmno6789',
    'vis'
);

-- Create session data
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

INSERT INTO user_sessions (user_id, session_name)
SELECT id, 'Get Revenge'
FROM users
WHERE email = 'john@example.com';

INSERT INTO user_sessions (user_id, session_name)
SELECT id, 'Master Will'
FROM users
WHERE email = 'vis@example.com';

-- Create an index
/* 
Indexes are patterns that allow for faster queries of the DB
*/
CREATE INDEX idx_user_sessions_user_id
ON user_sessions(user_id);

-- Analyze
EXPLAIN ANALYZE
SELECT *
FROM user_sessions
WHERE user_id = (
    SELECT id
    FROM users
    WHERE email = 'kelz@example.com'
);

-- Find a user by email
SELECT *
FROM users
WHERE email = 'kelz@example.com';

-- Find all sessions for a user
SELECT id, session_name
FROM user_sessions
WHERE user_id IN (SELECT id FROM users WHERE email = 'vis@example.com');

-- End a session
UPDATE user_sessions
SET ended_at = NOW()
WHERE session_name = 'Think in Computer';

-- Find all active sessions
SELECT *
FROM user_sessions
WHERE ended_at IS NULL;

-- Get nice rows of data
SELECT u.name, s.session_name, u.email, s.started_at
FROM user_sessions AS s
JOIN users AS u                                                        
ON s.user_id = u.id;

-- Find recently started sessions
SELECT *
FROM user_sessions
ORDER BY started_at DESC
LIMIT 10;
