-- +migrate Up
alter table markdown_posts
add foreign key (category_id) references categories(id);
alter table external_posts
add foreign key (category_id) references categories(id);
-- +migrate Down
alter table markdown_posts drop constraint foreign key (category_id) references categories(id);
alter table external_posts drop constraint foreign key (category_id) references categories(id);