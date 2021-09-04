-- +migrate Up
create table bookmarks (
  id int generated always as identity primary key,
  url text not null,
  description text not null,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null
);
-- +migrate Down
drop table if exists bookmarks;