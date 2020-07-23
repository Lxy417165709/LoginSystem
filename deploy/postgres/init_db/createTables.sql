drop table if exists tb_userAccountInformation CASCADE ;

create table tb_userAccountInformation(
    userId serial primary key,
    userEmail varchar(30) unique not null,
    userPassword varchar(50) not null,
    userType int not null,
    userRegisterTime bigint not null,
    userLastLoginTime bigint not null,
    salt	    varchar(10) not null,
    reserved2	varchar(1) not null
);
drop table if exists tb_userpersonalInformation CASCADE;
create table tb_userPersonalInformation(
    userId integer primary key references tb_userAccountInformation(userId) on delete cascade,
    userPhotoUrl varchar(80) not null,
    userName varchar(20) not null,
    userSex  integer not null,
    userContactPhone varchar(16) not null,
    userContactEmail varchar(30) not null,
    userBirthday bigint not null,
    reserved1	varchar(1) not null,
    reserved2	varchar(1) not null
);
