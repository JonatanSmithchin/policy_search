create table user
(
    UserID   int auto_increment
        primary key,
    UserName varchar(32) not null,
    Password varchar(32) not null,
    Age      smallint    null,
    Email    varchar(64) null,
    constraint user_UserName_uindex
        unique (UserName)
);

