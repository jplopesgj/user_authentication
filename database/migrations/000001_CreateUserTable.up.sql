CREATE SCHEMA IF NOT EXISTS user_auth;

CREATE TABLE IF NOT EXISTS user_auth.user (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(32),
    lastname    VARCHAR(32),
    username    VARCHAR(32) NOT NULL,
    email       VARCHAR(32) NOT NULL,
    password    VARCHAR(64) NOT NULL,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);