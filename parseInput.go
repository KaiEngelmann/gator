package main

import (
	"fmt"
)

func parseInput(args []string) (UserCommand, error) {
	if len(args) < 1 {
		return UserCommand{}, fmt.Errorf("command name required")
	}
	return UserCommand{name: args[0], args: args[1:]}, nil
}
