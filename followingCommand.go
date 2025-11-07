package main

import (
	"fmt"

	"github.com/kaiengelmann/gator/internal/database"
)

func following(s *AppState, cmd UserCommand, user database.User) error {
	if len(cmd.args) >= 1 {
		return fmt.Errorf("args not accepted")

	}
	feed_follows, err := s.db.GetFeedFollowsForUser(s.ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving feeds")
	}
	if len(feed_follows) == 0 {
		fmt.Println("User isn't following anyone yet")
	}
	for _, f := range feed_follows {
		fmt.Println("Feed Name: \n", f.FeedName)
	}
	return nil
}
