package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/kaiengelmann/gator/internal/config"
	"github.com/kaiengelmann/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	dbQueries := database.New(db)
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	appState := &AppState{
		db:  dbQueries,
		cfg: cfg,
		ctx: ctx,
	}
	cmds := &commandHandler{handlers: map[string]func(*AppState, UserCommand) error{}}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", reset)
	cmds.register("users", users)
	cmds.register("agg", agg)
	cmds.register("feeds", feeds)
	cmds.register("addfeed", middlewareLoggedIn(addfeed))
	cmds.register("follow", middlewareLoggedIn(follow))
	cmds.register("following", middlewareLoggedIn(following))
	cmds.register("unfollow", middlewareLoggedIn(unfollow))

	commandLine := os.Args
	if len(commandLine) < 2 {
		fmt.Print("please provide command")
		os.Exit(1)
	}
	cmd, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	if err := cmds.run(appState, cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
