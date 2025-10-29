package main

import (
	"errors"

	"github.com/kaiengelmann/gator/internal/config"
	"github.com/kaiengelmann/gator/internal/database"
)

type UserCommand struct {
	name string
	args []string
}

type AppState struct {
	db  *database.Queries
	cfg *config.Config
}

type commandHandler struct {
	handlers map[string]func(*AppState, UserCommand) error
}

func (u *commandHandler) run(as *AppState, cmd UserCommand) error {
	if command, ok := u.handlers[cmd.name]; ok {
		err := command(as, cmd)
		if err != nil {
			return err
		}
	} else {
		return errors.New("unknown command")
	}
	return nil
}
func (c *commandHandler) register(name string, f func(*AppState, UserCommand) error) {
	c.handlers[name] = f
}
