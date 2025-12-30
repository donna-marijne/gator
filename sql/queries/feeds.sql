-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, user_id, name, url)
values ($1, $2, $3, $4, $5, $6)
returning *;
