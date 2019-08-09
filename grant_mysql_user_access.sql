
-- winpty docker exec -it snippetbox_db bash
-- mysql -D snippetbox -u root -p

----------------------- MYSQL COMMANDS---------------------------------
-- CREATE USER 'web'@'localhost'
-- GRANT USAGE, SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';
-- ALTER USER 'web'@'localhost' IDENTIFIED BY 'sn1pp3tb0x';


CREATE USER 'web'@'%'
GRANT USAGE, SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'%';
ALTER USER 'web'@'%' IDENTIFIED BY 'sn1pp3tb0x';