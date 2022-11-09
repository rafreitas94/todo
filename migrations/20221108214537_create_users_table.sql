-- +goose Up
-- +goose StatementBegin
create table users (
  username text primary key,
  password text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
