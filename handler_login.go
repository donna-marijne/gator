package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		fmt.Printf("Usage: login <username>\n")
		return errors.New("missing arg")
	}

	userName := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no such user %s", userName)
		}
		return fmt.Errorf("failed to query user %s: %w", userName, err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Logged in as %s\n", user.Name)

	fmt.Printf("%+v\n", s.config)

	return nil
}
