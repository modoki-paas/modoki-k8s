CREATE TABLE users (
    seq INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    id VARCHAR(256) NOT NULL UNIQUE,
    type ENUM("user", "organization"),
    name VARCHAR(128) NOT NULL UNIQUE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    system_role ENUM("nothing", "admin")
);

ALTER TABLE users ADD INDEX name_index(name);