-- +migrate Up
create table articles (
    article_id integer not null primary key,
    title text,
    body text,
    created timestamp not null,
    updated timestamp
);

-- +migrate Down
DROP TABLE articles;
