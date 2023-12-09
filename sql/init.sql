CREATE DATABASE IF NOT EXISTS vault;

USE vault;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    lastname varchar(50) not null,
    genre varchar(50) not null,
    email varchar(50) not null unique,
    createdAt timestamp default current_timestamp()
);