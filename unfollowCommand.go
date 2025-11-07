package main

import (
	"fmt"

	"github.com/kaiengelmann/gator/internal/database"
)

func unfollow(s *AppState, cmd UserCommand, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("please provide URL")
	}
	feed_id, err := s.db.LookUpFeedsByUrl(s.ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("feed not found")
	}
	record := database.DeleteFollowRecordParams{
		UserID: user.ID,
		FeedID: feed_id.ID,
	}
	deleted, err := s.db.DeleteFollowRecord(s.ctx, record)
	if err != nil {
		return fmt.Errorf("failed to delete record %v", deleted.FeedID)
	}
	fmt.Printf("unfollowed %s", feed_id.Name)

	return nil
}
