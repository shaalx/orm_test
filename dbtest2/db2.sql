DROP DATABASE IF EXISTS os ;
CREATE database os charset=utf8;

use os;

CREATE TABLE IF NOT EXISTS user(
name char(20) PRIMARY KEY,
nickname char(20),
password char(64) NOT NULL,
age SMALLINT,
gentle TINYINT,
email char(30),
phone char(14),
address char(100)
)charset=utf8;


CREATE TABLE IF NOT EXISTS groups(
id INT PRIMARY KEY auto_increment,
name char(20) NOT NULL
)charset=utf8;

CREATE TABLE IF NOT EXISTS grel(
id INT PRIMARY KEY auto_increment,
gid INT NOT NULL,
uname char(20) NOT NULL,
nickname char(20),
FOERIGN KEY user(uname),
UNIQUE KEY (gid,uname)
)charset=utf8;