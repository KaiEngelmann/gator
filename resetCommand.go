package main

import (
	"context"
	"fmt"
	"os"
)

func reset(s *AppState, cmd UserCommand) error {
	ctx := context.Background()
	err := s.db.Reset(ctx)
	if err != nil {
		fmt.Printf("failed to reset\n")
		os.Exit(1)
	}
	fmt.Printf("reset complete\n")
	return nil
}
