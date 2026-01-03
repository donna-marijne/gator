package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/donnamarijne/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.args) > 0 {
		if parsed, err := strconv.ParseInt(cmd.args[0], 10, 32); err == nil {
			limit = int32(parsed)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%v\n%v\n\n%v\n\n---\n", post.Title.String, post.Url, post.Description.String)
		// fmt.Printf("%+v\n", post)
	}

	return nil
}
