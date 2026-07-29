-- Create a simple table
CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    display_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- projects references the id of the account
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES accounts(id),
    name TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Insert a new account
INSERT INTO accounts (email, display_name)
VALUES (
    'kelzbiz@example.com',
    'kelzbiz'
)
RETURNING id;

-- Insert a new project with the foreign key
INSERT INTO projects (owner_id, name)
VALUES (
    '33cd4f17-98ef-4343-83df-93bd38f59d03',
    'Learn Database Engineering'
);

-- The above was manual, this is better
INSERT INTO projects (owner_id, name)
SELECT id, 'Learn C#'
FROM accounts
WHERE email = 'kelzbiz@example.com';

-- drop the tables
DROP TABLE projects; -- this first since it is dependent on accounts
DROP TABLE accounts;

-- DROP DATABASE customer_db;