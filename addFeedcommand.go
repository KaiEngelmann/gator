package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kaiengelmann/gator/internal/database"
)

func addfeed(s *AppState, cmd UserCommand, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("please provide feed and URL")
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}
	new_feed, err := s.db.CreateFeed(s.ctx, feed)
	if err != nil {
		return err
	}
	feed_follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	if _, err := s.db.CreateFeedFollow(s.ctx, feed_follow); err != nil {
		return fmt.Errorf("error creating feed follow")
	}
	fmt.Printf("Feed created: %+v\n", new_feed)
	return nil
}
