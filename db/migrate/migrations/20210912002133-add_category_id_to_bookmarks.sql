-- +migrate Up
alter table bookmarks
add column category_id int;
-- +migrate Down
alter table bookmarks drop column category_id;