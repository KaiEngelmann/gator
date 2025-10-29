package main

import (
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
	appState := &AppState{
		db:  dbQueries,
		cfg: cfg,
	}

	cmds := &commandHandler{handlers: map[string]func(*AppState, UserCommand) error{}}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", reset)
	cmds.register("users", users)

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
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
