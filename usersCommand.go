package main

import (
	"context"
	"fmt"
	"os"
)

func users(s *AppState, cmd UserCommand) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		fmt.Printf("couldn't fetch users")
		os.Exit(1)
	}
	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}
