create table tendency
(
    TendencyID           int auto_increment
        primary key,
    UserID               int         not null,
    tendency_description varchar(32) not null,
    constraint user_desc
        unique (UserID, tendency_description),
    constraint UserID_FK
        foreign key (UserID) references user (UserID)
            on delete cascade
);

