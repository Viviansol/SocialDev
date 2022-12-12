
CREATE DATABASE IF NOT EXISTS socialDev;
USE socialDev;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nickName varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null unique,
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


insert into users(name, nickName, email, password)
values
("user1", "user_1", "user1@gmail.com","123456" ),
("user2", "user_2", "user1@gmail.com","123456" ),
("user3", "user_3", "user1@gmail.com","123456" ),
("user4", "user_4", "user1@gmail.com","123456" ),

    insert int followers(user_id, follower_id)
values
    (1,2),
    (2,3),
    (4,1)

CREATE TABLE publications(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    authorId int not null,
    FOREIGN KEY (authorId)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,
    createdAt timestamp default current_timestamp



)ENGINE = INNODB;

INSERT INTO publications( title, content, authorid) values ("title", "content", 1 );
INSERT INTO publications( title, content, authorid) values ("title2", "content2", 2 );
INSERT INTO publications( title, content, authorid) values ("title3", "content3", 3 );


