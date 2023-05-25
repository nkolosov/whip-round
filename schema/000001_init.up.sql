CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id         UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    login      VARCHAR(255) UNIQUE,
    email      VARCHAR(255) UNIQUE,
    birthdate  DATE,
    phone      VARCHAR(20) UNIQUE,
    balance    BIGINT,
    created_at TIMESTAMP
);

CREATE INDEX idx_users_login ON users (login);
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_phone ON users (phone);

CREATE TABLE refresh_sessions
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID         NOT NULL REFERENCES users (id),
    token      VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP    NOT NULL
);

CREATE INDEX idx_refresh_sessions_user_id ON refresh_sessions (user_id);
