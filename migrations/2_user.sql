-- +migrate Up
create table users (
    user_id integer not null primary key,
    name text not null,
    email text not null,
    salt text,
    salted text,
    created timestamp not null,
    updated timestamp
);

-- +migrate Down
DROP TABLE users;
