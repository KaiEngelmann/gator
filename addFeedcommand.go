package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kaiengelmann/gator/internal/database"
)

func addfeed(s *AppState, cmd UserCommand) error {
	if len(cmd.args) < 2 {
		fmt.Println("please provide feed and URL")
		os.Exit(1)
	}
	ctx := context.Background()
	current_user := s.cfg.CurrentUserName
	user, err := s.db.GetUser(ctx, current_user)
	if err != nil {
		fmt.Println("User doesn't Exist")
		os.Exit(1)
	}
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}
	new_feed, err := s.db.CreateFeed(ctx, feed)
	if err != nil {
		return err
	}
	fmt.Printf("Feed created: %+v\n", new_feed)
	return nil
}
