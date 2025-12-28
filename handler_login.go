package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		fmt.Printf("Usage: login <username>\n")
		return errors.New("missing arg")
	}

	userName := cmd.args[0]

	err := s.config.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("Set the user to %s\n", userName)

	fmt.Printf("%+v\n", s.config)

	return nil
}
