-- +goose Up
-- +goose StatementBegin
create extension pgcrypto;

create table users (
  username text primary key,
  password text not null
);

INSERT INTO users(username, password) VALUES
('usuario', crypt('senha', gen_salt('md5'))),
('usuario2', crypt('senha2', gen_salt('md5')));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
drop extension pgcrypto;
-- +goose StatementEnd
