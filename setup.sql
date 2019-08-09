CREATE DATABASE snippetbox;

USE snippetbox;

CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets  (title, content, created, expires) VALUES ('An old silent pond','An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Basho',UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY));

CREATE USER 'web'@'%'
GRANT USAGE, SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'%';
ALTER USER 'web'@'%' IDENTIFIED BY 'sn1pp3tb0x';