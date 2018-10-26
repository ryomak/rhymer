DROP DATABASE IF EXISTS rhymer_dev;
CREATE DATABASE rhymer_dev;

DROP USER IF EXISTS 'app' ;
FLUSH privileges;
CREATE USER 'app' IDENTIFIED BY 'password' ;
GRANT ALL ON rhymer_dev.* TO app@'localhost';
