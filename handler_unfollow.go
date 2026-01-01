package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/donnamarijne/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		fmt.Println("Usage: unfollow <url>")
		return errors.New("wrong number of args")
	}

	feedUrl := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to get feed by url: %w", err)
	}

	rowCount, err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete feed follow: %w", err)
	}
	if rowCount == 0 {
		fmt.Printf("%s is not following %s\n", user.Name, feedUrl)
		return nil
	}

	fmt.Printf("%s unfollowed %s\n", user.Name, feed.Name)

	return nil
}
