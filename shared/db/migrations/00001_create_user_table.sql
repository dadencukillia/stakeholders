-- +goose Up
CREATE TABLE users(
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  user_name VARCHAR(36) UNIQUE,
  full_name VARCHAR(200) DEFAULT '',
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS users;
