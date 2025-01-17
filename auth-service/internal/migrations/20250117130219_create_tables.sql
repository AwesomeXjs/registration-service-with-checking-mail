-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY ,
    email VARCHAR(255) NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL,
    role VARCHAR(5) NOT NULL CHECK (role IN ('user', 'admin')),
    verification BOOLEAN DEFAULT NULL,
    created_at timestamp not null default now(),
    updated_at timestamp
);
CREATE TABLE IF NOT EXISTS outbox (
    id SERIAL PRIMARY KEY,
    event VARCHAR(50) not null,
    payload VARCHAR(255) not null,
    created_at timestamp not null default now(),
    status VARCHAR(50) default 'pending'
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS outbox;