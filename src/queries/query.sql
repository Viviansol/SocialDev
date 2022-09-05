package queries


CREATE DATABASE IF NOT EXISTS socialDev;
USE socialDev;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nickName varchar(50) not null unique,
    email varchar(50) not null unique,

)