-- +goose Up
CREATE TABLE ${FILENAME}s
(
    id uuid not null primary key,
    foo varchar(255) not null,
    bar varchar(255),
    updated_at timestamp with time zone not null,
    created_at timestamp with time zone not null
);

-- +goose Down
DROP TABLE IF EXISTS ${FILENAME}s CASCADE;
