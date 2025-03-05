-- +goose Up
CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(16) not null,
    email VARCHAR not null,
    password VARCHAR not null,
    role VARCHAR not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

CREATE TABLE endpoints (
    id serial PRIMARY KEY,
    endpoint VARCHAR not null,
    role VARCHAR not null
);



-- +goose Down
DROP TABLE users;
DROP TABLE endpoints;
