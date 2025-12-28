package main

import (
	"fmt"
	"log"
	"os"

	"github.com/donnamarijne/gator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := state{
		config: &c,
	}

	cmds := NewCommands()
	cmds.register("login", handlerLogin)

	cmd, err := getCommand()
	if err != nil {
		fmt.Printf("Usage: gator <command>\n")
		os.Exit(1)
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Printf("Command %s failed: %v\n", cmd.name, err)
		os.Exit(1)
	}
}

func getCommand() (command, error) {
	var cmd command
	args := os.Args[1:]
	if len(args) < 1 {
		return cmd, fmt.Errorf("not enough args")
	}

	cmd = command{
		name: args[0],
		args: args[1:],
	}

	return cmd, nil
}
