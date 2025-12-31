-- name: CreateFeedFollow :one
insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
values ($1, $2, $3, $4, $5)
returning
	*,
	(select name from users where users.id = feed_follows.user_id) as user_name,
	(select name from feeds where feeds.id = feed_follows.feed_id) as feed_name;

-- name: GetFeedFollowsForUser :many
select sqlc.embed(feed_follows), sqlc.embed(feeds), sqlc.embed(users)
from feed_follows
	left join users on users.id = feed_follows.user_id
	left join feeds on feeds.id = feed_follows.feed_id
where users.id = $1;
