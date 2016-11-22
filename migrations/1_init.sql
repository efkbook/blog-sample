-- +migrate Up
create table articles (
    article_id integer not null primary key,
    title text,
    body text,
    created integer,
    updated integer
);

-- +migrate Down
DROP TABLE articles;
