-- +migrate Up
create TABLE external_posts (
  id int generated always as identity primary key,
  title text not null,
  url text not null,
  thumbnail_url text not null,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null
);
-- +migrate Down
drop table if exists external_posts;