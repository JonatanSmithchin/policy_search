create table admin_log
(
    id         bigint auto_increment
        primary key,
    admin_id   int default 114514 not null,
    admin_name varchar(255)       not null,
    method     varchar(10)        not null,
    url        varchar(150)       not null comment '请求的url',
    ip         varchar(30)        not null comment 'ip',
    code       varchar(10)        not null comment '返回的状态码',
    message    varchar(30)        not null comment '返回信息',
    created_at timestamp          null,
    updated_at timestamp          null
);

