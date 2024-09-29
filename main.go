package main

import (
	"time"

	"github.com/Ammar4372/pokedex/internal/pokeapi"
)

type command struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
	Pokemons map[string]pokeapi.Pokemon
}

func main() {
	cfg := Config{
		Client:   pokeapi.NewClient(10*time.Second, 5*time.Minute),
		Pokemons: make(map[string]pokeapi.Pokemon),
	}
	start_repl(cfg)
}
