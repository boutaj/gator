package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/boutaj/gator/internal/config"
	"github.com/boutaj/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cnf *config.Config
	db *database.Queries
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	dbURL := config.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	
	gstate := &state{
		cnf: config,
		db: database.New(db),
	}

	commands := &commands{
		handler: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

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