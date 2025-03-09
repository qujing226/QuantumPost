CREATE DATABASE sms_mail;
USE sms_mail;

create table biz (
    id int not null auto_increment,

    name      varchar(255) not null,
    applicant varchar(255) not null,
    applicantId int not null,
    status      int not null default 0,

    sms_cap int not null,
    mail_cap int not null,
    sms_num int,
    mail_num int,
    primary key (id)
);

create table sms (
    id int not null auto_increment,
    bizId int not null,
    phone varchar(255) not null,
    content varchar(255) not null,


    status int not null default 0,
    primary key (id)
);
