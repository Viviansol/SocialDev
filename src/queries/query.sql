
CREATE DATABASE IF NOT EXISTS socialDev;
USE socialDev;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nickName varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(20) not null unique,
    createdAt timestamp default  current_timestamp()


)ENGINE = INNODB;


CREATE TABLE followers(
    user_id int not null ,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    PRIMARY KEY (user_id, follower_id)

)ENGINE = INNODB;