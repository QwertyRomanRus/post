-- +goose Up
-- +goose StatementBegin
CREATE table posts
(
    id         serial,
    title      VARCHAR(255) not null,
    content    TEXT not null,
    user_id    INTEGER      not null,
    rating     INTEGER      not null,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),

    primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table posts;
-- +goose StatementEnd
