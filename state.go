package main

import (
	"github.com/donna-marijne/gator/internal/config"
	"github.com/donna-marijne/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
