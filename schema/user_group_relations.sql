CREATE TABLE user_group_relations (
    id INTEGER NOT NULL PRIMARY KEY,
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    permission JSON
);