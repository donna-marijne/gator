package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/donnamarijne/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		userName := s.config.CurrentUserName
		user, err := s.db.GetUser(context.Background(), userName)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("no such user: %s", userName)
			}

			return fmt.Errorf("failed to query user %s: %w", userName, err)
		}

		return handler(s, cmd, user)
	}
}
