CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    balance FLOAT NOT NULL CHECK (balance >= 0)
);