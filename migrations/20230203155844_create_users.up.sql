CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null,
    email varchar not null unique
)