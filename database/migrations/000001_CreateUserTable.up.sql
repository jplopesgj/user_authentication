CREATE SCHEMA finance;

CREATE TABLE finance.user (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50),
    lastname    VARCHAR(50),
    password    VARCHAR(50),
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);