CREATE TABLE apps (
    seq INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    id VARCHAR(128) NOT NULL UNIQUE,
    owner VARCHAR(128),
    name VARCHAR(128) NOT NULL UNIQUE,
    spec JSON,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);

ALTER TABLE apps ADD INDEX name_index(name);