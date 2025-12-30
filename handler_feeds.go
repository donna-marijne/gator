package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s: %s (%s)\n", feed.Feed.Name, feed.Feed.Url, feed.User.Name)
	}

	return nil
}
