package main

import (
	"context"
	"fmt"

	"github.com/boutaj/gator/internal/rss"
)

const URL = "https://www.wagslane.dev/index.xml"

func handlerRss(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), URL)
	if err != nil {
		return err
	}
	fmt.Print(feed)
	return nil
}