DROP DATABASE IF EXISTS dbt;
CREATE DATABASE dbt charset=utf8;
use dbt;
CREATE TABLE user(
id int PRIMARY KEY auto_increment,
name char(20),
pid int,
UNIQUE KEY(name)
)charset=utf8;

CREATE TABLE profile(
id int PRIMARY KEY auto_increment,
name char(50),
uid int
)charset=utf8;

CREATE TABLE post(
id int PRIMARY KEY auto_increment,
uid int,
tsid int
)charset=utf8;

CREATE TABLE tag(
id int PRIMARY KEY auto_increment,
psid int
)charset=utf8;