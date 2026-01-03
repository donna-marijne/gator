-- +goose Up
create table posts (
	id uuid primary key,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	url text unique not null,
	title text,
	description text,
	published_at timestamp,
	feed_id uuid not null references feeds on delete cascade
);

-- +goose Down
drop table posts;

