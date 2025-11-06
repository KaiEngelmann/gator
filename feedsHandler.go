package main

import (
	"context"
	"fmt"
	"os"
)

func feeds(s *AppState, cmd UserCommand) error {
	ctx := context.Background()
	all_feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("All Feeds: %+v\n", all_feeds)

	return nil
}
