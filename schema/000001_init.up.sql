CREATE TABLE users_token
(
    id serial not null unique,
    guid varchar(36) not null,
    email varchar(255) not null,
    ip varchar(255) not null,
    refresh_token varchar(255) not null
);