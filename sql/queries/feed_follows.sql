-- name: CreateFeedFollow :one
insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
values ($1, $2, $3, $4, $5)
returning
	*,
	(select name from users where users.id = feed_follows.user_id) as user_name,
	(select name from feeds where feeds.id = feed_follows.feed_id) as feed_name;
