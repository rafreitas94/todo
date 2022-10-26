-- +goose Up
-- +goose StatementBegin
create table tasks(
  id text primary key,
  subject text not null,
  description text not null,
  status text not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
