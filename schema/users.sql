CREATE TABLE users (
    seq INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    type ENUM("user", "organization"),
    id VARCHAR(256) NOT NULL UNIQUE,
    name VARCHAR(128) NOT NULL UNIQUE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    system_role ENUM("nothing", "admin")
);

ALTER TABLE users ADD INDEX name_index(name);