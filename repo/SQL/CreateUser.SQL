/*
 Source Server         : 本地
 Source Server Type    : PostgreSQL
 Source Host           : localhost:5432
 Source Schema         : carshop
*/
create user bf with password 'bf123'
create database bfdb owner bf;
grant all privileges on database bfdb to bf;



