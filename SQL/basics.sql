-- create the DB with CREATE DATABASE customer_db;
-- then connect with \c customer_db;
CREATE TABLE customers (
    id UUID PRIMARY KEY, 
    name TEXT NOT NULL, 
    email TEXT NOT NULL UNIQUE
);

-- insert into the DB
INSERT INTO customers (id, name, email)
VALUES (
    gen_random_uuid(),
    'Ada',
    'ada@example.com'
);

-- update an entry, when you are dumb and enter the data wrong
UPDATE customers
SET name = 'Grace'
WHERE email = 'grace@example.com';

-- remove an entry
DELETE from customers
WHERE email = 'kelz@example.com';

-- basic query
SELECT * from customers
WHERE name = "Kelz";

-- drop the entire table, destructive obviously
DROP TABLE customers;

-- delete the DB when finished with:
-- DROP DATABASE customer_db;