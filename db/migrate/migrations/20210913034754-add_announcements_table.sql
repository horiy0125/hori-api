-- +migrate Up
create table announcements (
  id int generated always as identity primary key,
  content text not null,
  url text not null,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null
);
-- +migrate Down
drop table if exists announcements;