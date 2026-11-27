package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/boutaj/gator/internal/database"
)

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("the follow handler expects a single argument, the feed url.")
	}
	
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])

	params := &database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: int32(feed.ID),
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), *params)
	if err != nil {
		return err
	}

	fmt.Printf("%s - %s\n", feedFollow.UserName, feedFollow.FeedName)

	return  nil
}

func handleFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("%s\n", feedFollow.Name)
	}

	return  nil
}

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("the unfollow handler expects a single argument, the feed url.")
	}
	
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])

	params := &database.DeleteFeedFollowForUserParams{
		UserID: user.ID,
		FeedID: int32(feed.ID),
	}

	err = s.db.DeleteFeedFollowForUser(context.Background(), *params)
	if err != nil {
		return err
	}

	return  nil
}