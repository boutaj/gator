package main

import (
	"log"
	"os"

	"github.com/boutaj/gator/internal/config"
)

type state struct {
	cnf *config.Config
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	gstate := &state{
		cnf: config,
	}

	commands := &commands{
		handler: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("expected at least 1 argument, got 0")
	}

	if err := commands.run(gstate, command{
		name: arguments[1],
		args: arguments[2:],
	}); err != nil {
		log.Fatal(err)
	}
}