drop database if exists qq;
create database qq charset=utf8;

use qq;

create table user(
id int primary key auto_increment,
name char(20),
profile_id int,
unique key(name),
key profile(profile_id)
)charset=utf8;

create table profile(
id int primary key auto_increment,
user_id int,
key user(user_id)
)charset=utf8;


create table userz(
id int primary key auto_increment,
name char(20),
unique key(name),
friends_id int
)charset=utf8;

create table friends(
id int primary key auto_increment,
userz_id int,
userzt_id int,
key userz(userz_id,userzt_id)
)
