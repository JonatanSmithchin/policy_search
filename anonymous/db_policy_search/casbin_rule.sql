create table casbin_rule
(
    id    bigint auto_increment
        primary key,
    ptype varchar(100) null,
    v0    varchar(100) null,
    v1    varchar(100) null,
    v2    varchar(100) null,
    v3    varchar(100) null,
    v4    varchar(100) null,
    v5    varchar(100) null,
    constraint idx_casbin_rule
        unique (ptype, v0, v1, v2, v3, v4, v5)
)
    comment '权限规则表';

