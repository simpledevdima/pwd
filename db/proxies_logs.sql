create table if not exists proxies_logs
(
    id        int auto_increment comment 'Record ID'
        primary key,
    dt_create datetime default CURRENT_TIMESTAMP not null comment 'Date and time the entry was added',
    id_proxy  int                                not null comment 'Proxy ID',
    url       varchar(255)                       not null comment 'page address to open',
    domain    varchar(100)                       not null comment 'page domain',
    code      mediumint                          not null comment 'HTTP code returned by the server',
    duration  float(5, 2)                        null comment 'Total request time'
)
    comment 'proxy usage log';
