-- +migrate Up
alter table markdown_posts
add column publish boolean not null default false;
-- +migrate Down
alter table markdown_posts drop column publish;