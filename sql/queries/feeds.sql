-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, user_id, name, url)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
select sqlc.embed(feeds), sqlc.embed(users)
from feeds 
	left join users on feeds.user_id = users.id;

-- name: GetFeedByUrl :one
select *
from feeds
where url = $1;

-- name: GetNextFeedToFetch :one
select *
from feeds
order by last_fetched_at asc nulls first
limit 1;

-- name: MarkFeedFetched :one
update feeds
set last_fetched_at = current_timestamp, updated_at = current_timestamp
where id = $1
returning *;

