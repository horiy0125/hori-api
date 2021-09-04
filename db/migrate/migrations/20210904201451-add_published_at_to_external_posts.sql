-- +migrate Up
alter table external_posts
add column published_at timestamptz not null;
-- +migrate Down
alter table external_posts drop column published_at;