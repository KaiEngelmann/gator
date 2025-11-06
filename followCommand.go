package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kaiengelmann/gator/internal/database"
)

func follow(s *AppState, cmd UserCommand) error {
	if len(cmd.args) != 1 {
		fmt.Printf("Please provide URL")
		os.Exit(1)
	}
	ctx := context.Background()
	current_user := s.cfg.CurrentUserName
	user, err := s.db.GetUser(ctx, current_user)
	if err != nil {
		fmt.Println("User doesn't Exist")
		os.Exit(1)
	}
	feed, err := s.db.LookUpFeedsByUrl(ctx, cmd.args[0])
	if err != nil {
		fmt.Println("Feed not found")
		os.Exit(1)
	}
	feed_follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	new_feed, err := s.db.CreateFeedFollow(ctx, feed_follow)
	if err != nil {
		fmt.Println("Error creating feed follow")
		os.Exit(1)
	}
	fmt.Printf("%v %v", new_feed.FeedName, current_user)
	return nil
}
