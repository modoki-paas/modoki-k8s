CREATE TABLE role_bindings (
    seq INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_seq INTEGER NOT NULL,
    target_seq INTEGER NOT NULL,
    role_name VARCHAR(128)
);

ALTER TABLE role_bindings ADD UNIQUE INDEX user_target_index(user_seq, target_seq);
ALTER TABLE role_bindings ADD INDEX user_index(user_seq);
ALTER TABLE role_bindings ADD INDEX target_index(target_seq);
