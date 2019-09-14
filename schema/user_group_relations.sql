CREATE TABLE user_group_relations (
    id INTEGER NOT NULL PRIMARY KEY,
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    permission JSON
);

ALTER TABLE services ADD INDEX group_user_index(group_id, user_id);