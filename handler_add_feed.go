package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/donnamarijne/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		fmt.Println("Usage: addfeed <name> <url>")
		return errors.New("incorrect number of args")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	userName := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("failed to query user %s: %w", userName, err)
	}

	now := time.Now()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("failed to create the feed %q (%s): %w", name, url, err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow for %s: %w", url, err)
	}

	fmt.Printf("Created: %+v\n", feed)

	return nil
}
