create table if not exists user_agents
(
    id   int auto_increment comment 'user agent id'
        primary key,
    name varchar(255) null comment 'user agent name'
)
    comment 'list of popular user agents';
