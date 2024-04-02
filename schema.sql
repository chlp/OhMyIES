create table feeds
(
    id            varchar(32)  not null,
    name          varchar(64)  not null,
    rss_url       varchar(128) not null,
    rss_last_id   varchar(128) not null,
    rss_last_dt   datetime,
    created_at_dt datetime     not null default NOW,
    constraint feeds_pk
        primary key (id)
);

create index feeds_rss_url_idx
    on feeds (rss_url);

create table chats
(
    id             varchar(32)  not null,
    feed_id        varchar(32)  not null,
    name           varchar(64)  not null,
    filter         int          not null,
    chat_id        varchar(128) not null,
    rss_last_id    varchar(128) not null,
    last_msg_dt    datetime,
    last_msg_error varchar(256) not null,
    msg_count      int          not null,
    created_at_dt  datetime     not null default NOW,
    constraint chats_pk
        primary key (id)
);
