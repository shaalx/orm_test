DROP DATABASE IF EXISTS os ;
CREATE database os charset=utf8;

use os;

CREATE TABLE IF NOT EXISTS user(
id INT PRIMARY KEY auto_increment,
name char(20) NOT NULL,
nickname char(20),
password char(64) NOT NULL,
age SMALLINT,
gentle TINYINT,
email char(30),
phone char(14),
address char(100),
UNIQUE KEY (name)
)charset=utf8;


CREATE TABLE IF NOT EXISTS groups(
id INT PRIMARY KEY auto_increment,
name char(20) NOT NULL
)charset=utf8;

CREATE TABLE IF NOT EXISTS grel(
id INT PRIMARY KEY auto_increment,
gid INT NOT NULL,
uid INT NOT NULL,
nickname char(20),
KEY user(uid),
KEY groups(gid),
UNIQUE KEY(gid,uid)
)charset=utf8;

CREATE TABLE IF NOT EXISTS message(
id INT PRIMARY KEY auto_increment,
uin INT NOT NULL,
uout INT NOT NULL,
content BLOB NOT NULL,
timestamp INT NOT NULL,
KEY user(uin,uout)
)charset=utf8;

CREATE TABLE IF NOT EXISTS item(
id INT PRIMARY KEY auto_increment,
uid INT NOT NULL,
name char(20) not null,
KEY user(uid)
)charset=utf8;

CREATE TABLE IF NOT EXISTS friends(
id INT PRIMARY KEY auto_increment,
uid INT NOT NULL,
iid INT NOT NULL,
fuid INT NOT NULL,
KEY user(uid,fuid),
KEY item(iid),
UNIQUE KEY(iid,fuid)
)charset=utf8;

CREATE TABLE IF NOT EXISTS login(
id INT PRIMARY KEY auto_increment,
uid INT NOT NULL,
timestamp INT NOT NULL,
ip CHAR(15),
KEY user(uid)
)charset=utf8;