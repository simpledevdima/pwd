create table if not exists proxies
(
    id           int auto_increment comment 'Proxy server address ID'
        primary key,
    dt_create    datetime             default CURRENT_TIMESTAMP not null comment 'Date and time the server address was added to the database',
    dt_last_used datetime                                       null comment 'Date and time the server was last used',
    working      enum ('U', 'Y', 'N') default 'U'               null comment 'Is the proxy server working? U - Unknown, Y - Yes, N - No',
    ip           varchar(15)                                    not null comment 'IP address of the proxy server',
    port         mediumint                                      not null comment 'Proxy server port',
    type         enum ('ssl', 'http')                           not null comment 'Proxy server type'
)
    comment 'Table with proxy server addresses';
