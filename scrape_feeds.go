package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/donna-marijne/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("failed to get next feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark the feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch the RSS feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Url:         item.Link,
			Title:       toNullString(item.Title.String()),
			Description: toNullString(item.Description.String()),
			PublishedAt: toNullTime(item.PubDate),
			FeedID:      feed.ID,
		})
		if err != nil {
			if pqerr, ok := err.(*pq.Error); ok {
				if pqerr.Code == "23505" {
					continue
				}
			}

			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("New post: %v (%v)\n", post.Title, post.Url)
	}

	return nil
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}

	return sql.NullString{String: s, Valid: true}
}

func toNullTime(s string) sql.NullTime {
	t, err := parseTime(s)
	if err != nil {
		fmt.Printf("Failed to parse time: %v\n", err)
		return sql.NullTime{Valid: false}
	}

	return sql.NullTime{Time: t, Valid: true}
}
