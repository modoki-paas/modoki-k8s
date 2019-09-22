CREATE TABLE services (
    id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    owner INTEGER,
    name VARCHAR(128) NOT NULL UNIQUE,
    config JSON,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);

ALTER TABLE services ADD INDEX name_index(name);