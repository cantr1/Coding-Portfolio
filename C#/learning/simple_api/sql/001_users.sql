-- Create the `users` table
CREATE TABLE users (
    id uuid NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE
);

-- Insert data into the `users` table
INSERT INTO users (id, name, username, email) 
VALUES ('Kelly Cantrell', 'kelz', 'kelz@example.com');

-- Drop the `users` table
DROP TABLE users;