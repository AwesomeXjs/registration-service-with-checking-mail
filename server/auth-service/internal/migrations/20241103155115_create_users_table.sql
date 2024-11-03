-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    email VARCHAR(255) NOT NULL,
    hash_password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
DROP TABLE IF EXISTS users
