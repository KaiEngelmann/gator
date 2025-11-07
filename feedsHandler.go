package main

import (
	"fmt"
	"os"
)

func feeds(s *AppState, cmd UserCommand) error {
	all_feeds, err := s.db.GetFeeds(s.ctx)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("All Feeds: %+v\n", all_feeds)

	return nil
}
