CREATE TABLE user_group_relations (
    id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    permission JSON
);

ALTER TABLE user_group_relations ADD INDEX group_user_index(group_id, user_id);
ALTER TABLE user_group_relations ADD INDEX group_index(group_id);
ALTER TABLE user_group_relations ADD INDEX user_index(user_id);
