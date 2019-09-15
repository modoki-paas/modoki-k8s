CREATE TABLE tokens (
    id INTEGER NOT NULL PRIMARY KEY,
    token VARCHAR(512) NOT NULL UNIQUE,
    organization INTEGER NOT NULL,
    author INTEGER NOT NULL,
    permission JSON,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
);

ALTER TABLE tokens ADD INDEX author_index(author);
ALTER TABLE tokens ADD INDEX organization_index(organization);
