package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/donna-marijne/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		fmt.Printf("Usage: register <username>\n")
		return errors.New("missing arg")
	}

	userName := cmd.args[0]

	now := time.Now()
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      userName,
	})
	if err != nil {
		return fmt.Errorf("CreateUser with Name = %s failed: %w", userName, err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s registered\n", user.Name)
	fmt.Printf("%+v\n", user)

	return nil
}
