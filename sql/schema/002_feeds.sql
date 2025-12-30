-- +goose Up
create table feeds (
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	user_id uuid not null references users on delete cascade,
	name text not null,
	url text unique not null
);

-- +goose Down
drop table feeds;

