create table read_record
(
    RecID         int auto_increment
        primary key,
    UserID        int           not null,
    IP            varchar(32)   not null,
    Read_time     datetime      not null,
    content       varchar(64)   not null,
    Read_duration int default 0 not null,
    tag           varchar(32)   not null,
    constraint record_UserID_FK
        foreign key (UserID) references user (UserID)
            on delete cascade
);

