package main

import (
	"fmt"
	"os"

	"github.com/kaiengelmann/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *AppState, cmd UserCommand, user database.User) error) func(*AppState, UserCommand) error {
	return func(s *AppState, cmd UserCommand) error {
		currentUserName := s.cfg.CurrentUserName
		if currentUserName == "" {
			fmt.Println("not logged in")
			os.Exit(1)
		}
		user, err := s.db.GetUser(s.ctx, currentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)

	}
}
