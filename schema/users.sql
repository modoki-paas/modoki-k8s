CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL UNIQUE,
    auth_method VARCHAR(64) NOT NULL DEFAULT "password",
    password VARCHAR(512),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
);