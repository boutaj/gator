package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/boutaj/gator/internal/database"
	"github.com/boutaj/gator/internal/rss"
)

const URL = "https://www.wagslane.dev/index.xml"

func handleRss(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), URL)
	if err != nil {
		return err
	}
	fmt.Print(feed)
	return nil
}

func handleFeed(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("the login handler expects two argument, the name and url.")
	}

	user, err := s.db.GetUser(context.Background(), s.cnf.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.args[0]
	url  := cmd.args[1]

	params := &database.CreateFeedParams{
		Name: name,
		Url: url,
		UserID: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feed, err := s.db.CreateFeed(context.Background(), *params)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)

	return nil
}

func handleFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feeds)

	return  nil
}