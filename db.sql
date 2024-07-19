create database if not exists belajar_golang_restful_api;
create database if not exists belajar_golang_restful_api_test;

use belajar_golang_restful_api;
use belajar_golang_restful_api_test;

create table categories (
    id int primary key auto_increment,
    name varchar(200) not null
)engine innoDB;