-- +migrate Up
create table articles (
    article_id integer not null primary key,
    title text not null,
    body text not null,
    user_id integer not null,
    created timestamp not null,
    updated timestamp
);

-- +migrate Down
DROP TABLE articles;
