CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создаем таблицу users
CREATE TABLE users
(
    id         UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    login      VARCHAR(255),
    email      VARCHAR(255),
    birthdate  DATE,
    phone      VARCHAR(20),
    balance    BIGINT,
    created_at TIMESTAMP
);

CREATE TABLE refresh_sessions
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID         NOT NULL REFERENCES users (id),
    token      VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP    NOT NULL
);
