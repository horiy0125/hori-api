-- +migrate Up
alter table markdown_posts
alter column category_id
set not null;
alter table external_posts
alter column category_id
set not null;
alter table bookmarks
alter column category_id
set not null;
-- +migrate Down
alter table markdown_posts
alter column category_id drop not null;
alter table external_posts
alter column category_id drop not null;
alter table bookmarks
alter column category_id drop not null;