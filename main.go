package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/donna-marijne/gator/internal/config"
	"github.com/donna-marijne/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", c.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := state{
		config: &c,
		db:     dbQueries,
	}

	cmds := NewCommands()

	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerFeeds)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

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
