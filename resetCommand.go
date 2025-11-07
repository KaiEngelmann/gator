package main

import (
	"fmt"
	"os"
)

func reset(s *AppState, cmd UserCommand) error {
	err := s.db.Reset(s.ctx)
	if err != nil {
		fmt.Printf("failed to reset\n")
		os.Exit(1)
	}
	fmt.Printf("reset complete\n")
	return nil
}
