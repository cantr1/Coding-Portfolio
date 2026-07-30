-- A better schema, so that creation is more simple
CREATE TABLE customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    name TEXT NOT NULL, 
    email TEXT NOT NULL, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- the insert now only requires name and email
INSERT INTO customers (name, email)
VALUES (
    'Kelly',
    'kelz@example.com'
),
(
    'Ada',
    'ada@example.com"'
),
(
    'Grace',
    'grace@example.com'
);
