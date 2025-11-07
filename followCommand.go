package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kaiengelmann/gator/internal/database"
)

func follow(s *AppState, cmd UserCommand, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("please provide URL")
	}
	feed, err := s.db.LookUpFeedsByUrl(s.ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("feed not found")
	}
	feed_follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	new_feed, err := s.db.CreateFeedFollow(s.ctx, feed_follow)
	if err != nil {
		return fmt.Errorf("error creating feed follow")
	}
	fmt.Printf("%v %v", new_feed.FeedName, user.Name)
	return nil
}
