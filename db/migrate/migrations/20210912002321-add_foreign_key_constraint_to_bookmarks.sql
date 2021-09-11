-- +migrate Up
alter table bookmarks
add foreign key (category_id) references categories(id);
-- +migrate Down
alter table bookmarks drop constraint foreign key (category_id) references categories(id);