-- +goose Up
-- +goose StatementBegin
create table users (
  username text primary key,
  password text not null
);

INSERT INTO users(username, password) VALUES
('usuario', 'senha'),
('usuario2', 'senha2');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
