package main

import (
	"fmt"
)

func agg(s *AppState, cmd UserCommand) error {
	url := "https://www.wagslane.dev/index.xml"
	resp, err := fetchFeed(s.ctx, url)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}
