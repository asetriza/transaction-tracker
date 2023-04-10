CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  state TEXT NOT NULL,
  amount FLOAT NOT NULL,
  transaction_id TEXT NOT NULL
);