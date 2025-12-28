package main

import "fmt"

type commands struct {
	handlers map[string]func(*state, command) error
}

func NewCommands() commands {
	return commands{
		handlers: map[string]func(*state, command) error{},
	}
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("no such command: %s", cmd.name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
