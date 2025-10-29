package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kaiengelmann/gator/internal/database"
)

func handlerRegister(s *AppState, cmd UserCommand) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("please provide username")
	}
	userId := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	cxt := context.Background()
	_, err := s.db.GetUser(cxt, userId.Name)
	if err == nil {
		fmt.Println("User already exists")
		os.Exit(1)
	}
	newUser, err := s.db.CreateUser(cxt, userId)
	if err != nil {
		return err
	}
	s.cfg.SetUser(userId.Name)
	fmt.Printf("User created: %+v\n", newUser)

	return nil
}
