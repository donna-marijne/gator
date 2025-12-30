package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       URLEncodedString `xml:"title"`
		Link        string           `xml:"link"`
		Description URLEncodedString `xml:"description"`
		Item        []RSSItem        `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       URLEncodedString `xml:"title"`
	Link        string           `xml:"link"`
	Description URLEncodedString `xml:"description"`
	PubDate     string           `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %w", err)
	}

	req.Header.Set("user-agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send a request: %w", err)
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read a response body: %w", err)
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(buf, &rssFeed)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML from request body: %w", err)
	}

	return &rssFeed, nil
}
