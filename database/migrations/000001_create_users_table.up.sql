CREATE TABLE IF NOT EXISTS users
(
    id         uuid PRIMARY KEY,
    first_name text,
    last_name  text,
    nickname   text,
    email      text      NOT NULL,
    country    text,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    CONSTRAINT UNIQUE (email)
);