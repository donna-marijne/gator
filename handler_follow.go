package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/donnamarijne/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		fmt.Println("Usage: follow <url>")
		return errors.New("wrong number of args")
	}

	feedUrl := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to get feed by URL %s: %w", feedUrl, err)
	}

	now := time.Now()
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create follow: %w", err)
	}

	fmt.Printf("%s followed %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
