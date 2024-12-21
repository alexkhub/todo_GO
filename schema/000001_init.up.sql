CREATE Table users(
id SERIAL primary key ,
username varchar(150),
email varchar(150) NULL,
hash_password varchar(150)

);

Create Table tasks(
id SERIAL primary key ,
user_id int REFERENCES users (id),
title varchar(100),
description text NULL,
status varchar(20) check(status = 'not completed' or status = 'completed' or status= 'in_progress' ) default 'not completed',
datetime_create timestamp

);