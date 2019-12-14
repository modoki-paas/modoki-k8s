CREATE TABLE user_org_relations (
    seq INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    org_seq INTEGER NOT NULL,
    user_seq INTEGER NOT NULL
);

ALTER TABLE user_org_relations ADD INDEX org_user_index(org_seq, user_seq);
ALTER TABLE user_org_relations ADD INDEX org_index(org_seq);
ALTER TABLE user_org_relations ADD INDEX user_index(user_seq);
