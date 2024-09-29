package main

import (
	"time"

	"github.com/Ammar4372/pokedex/internal/pokeapi"
)

type command struct {
	name        string
	description string
	callback    func(*Config, string) error
}

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func main() {
	cfg := Config{
		Client: pokeapi.NewClient(3*time.Second, 5*time.Minute),
	}
	start_repl(cfg)
}
