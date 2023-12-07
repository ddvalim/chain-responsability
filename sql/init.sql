CREATE DATABASE IF NOT EXISTS p;

USE p;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    createdAt timestamp default current_timestamp()
);