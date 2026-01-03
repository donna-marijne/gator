-- name: CreatePost :one
insert into posts (id, url, title, description, published_at, feed_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetPostsForUser :many
select posts.*
from posts
	left join feed_follows on feed_follows.feed_id = posts.feed_id
where feed_follows.user_id = $1
order by posts.published_at desc
limit $2;

