package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/donnamarijne/gator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	userName, err := getUserName()
	if err != nil {
		log.Fatal(err)
	}

	err = c.SetUser(userName)
	if err != nil {
		log.Fatal(err)
	}

	c2, err := config.Read()
	fmt.Println(c2)
}

func getUserName() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get user info from the os: %w", err)
	}

	return u.Username, nil
}
