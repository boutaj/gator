package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	username := cmd.args[0]
	if err := s.cnf.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("The username %s has been set!", username)

	return nil
}