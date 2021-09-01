create TABLE markdown_posts (
  id int generated always as identity primary key,
  title text not null,
  body text not null,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null
);
-- create TABLE external_posts (
--   id int generated always as identity primary key,
--   title text not null,
--   url text not null,
--   thumbnail_url text not null,
--   created_at timestamptz not null default current_timestamp,
--   updated_at timestamptz not null
-- );