-- +migrate Up
create table users (
    user_id integer not null primary key,
    name text,
    email text,
    salt text,
    salted text,
    created integer,
    updated integer
);

-- +migrate Down
DROP TABLE users;
