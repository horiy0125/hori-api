-- +migrate Up
create TABLE categories (
  id int generated always as identity primary key,
  name text not null,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null
);
-- +migrate Down
drop table if exists categories;