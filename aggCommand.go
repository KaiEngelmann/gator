package main

import (
	"context"
	"fmt"
)

func agg(s *AppState, cmd UserCommand) error {
	url := "https://www.wagslane.dev/index.xml"
	cxt := context.Background()
	resp, err := fetchFeed(cxt, url)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}
