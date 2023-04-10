CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  transaction_id VARCHAR NOT NULL,
  account_id INT NOT NULL REFERENCES accounts(id),
  state INT NOT NULL,
  amount FLOAT NOT NULL,
  is_canceled BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL
);