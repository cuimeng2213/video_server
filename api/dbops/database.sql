create database video_server

Table: users -- 用户表

id unsigened int, primary_key, auto_increment
login_name varchar(64), unique key,
pwd Text
 create table users (id int(10) unsigned primary key auto_increment , login_name varchar(64), pwd text);

Table: video_info
id varchar(64), primary_key, not null --其中存放 uuid
author_id unsigened int 
name text --视频名字
display_ctime TEXT --显示在页面得名字
create_time DATETIME -- 创建时间
create table video_info(id varchar(64), author_id int(10), name text, display_ctime text, create_time datetime);

TABLE: comments
id varchar(64), primary_key, not null -- 评论数比较多
video_id varchar(64)
author_id unsigened int 
content TEXT 
time DATETIME
create table comments (id varchar(64), video_id varchar(64), author_id int(10), content text, time datetime);
-- 保存用户登陆信息
TABLE sessions
session_id TINYTEXT,primary_key,not null
TTL TINYTEXT -- 过期时间
login_name varchar(64)
 create table sessions(session_id tinytext, TTL tinytext, login_name text);
