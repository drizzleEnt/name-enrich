CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    age int not null,
    gender varchar(255) not null,
    nationality varchar(255) not null
);