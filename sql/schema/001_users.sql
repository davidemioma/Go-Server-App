-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL,
    username TEXT NOT NULL,
    role TEXT DEFAULT 'user' CHECK (role IN ('user', 'admin')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;