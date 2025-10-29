package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *AppState, cmd UserCommand) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("please provide username")
	}
	cxt := context.Background()
	_, err := s.db.GetUser(cxt, cmd.args[0])
	if err != nil {
		fmt.Println("User doesn't Exist")
		os.Exit(1)
	}
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("User: %s has been set\n", s.cfg.CurrentUserName)
	return nil
}
