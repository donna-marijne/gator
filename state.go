package main

import (
	"github.com/donnamarijne/gator/internal/config"
	"github.com/donnamarijne/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
