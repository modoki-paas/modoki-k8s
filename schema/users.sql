CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL UNIQUE,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
);