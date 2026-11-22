package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/boutaj/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	username := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	if err := s.cnf.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("The username %s has been set!\n", user.Name)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	params := &database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), *params)
	if err != nil {
		return err
	}

	if err := s.cnf.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("The username %s has been registred!\n", user.Name)

	return nil
}