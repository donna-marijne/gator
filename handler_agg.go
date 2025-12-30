package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feedUrl := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
