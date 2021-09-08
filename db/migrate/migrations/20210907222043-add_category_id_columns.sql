-- +migrate Up
alter table markdown_posts
add column category_id int;
alter table external_posts
add column category_id int;
-- +migrate Down
alter table external_posts drop column category_id;
alter table markdown_posts drop column category_id;