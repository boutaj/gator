package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handler map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.handler[cmd.name]
	if !ok {
		return fmt.Errorf("command %s does not exist!", cmd.name)
	}
	return fn(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handler[name] = f
}