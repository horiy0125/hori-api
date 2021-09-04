-- +migrate Up
create table markdown_posts (
  id int generated always as identity primary key,
  title text not null,
  body text not null,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null
);
-- +migrate Down
drop table if exists markdown_posts;