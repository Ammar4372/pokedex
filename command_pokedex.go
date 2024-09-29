package main

import "fmt"

func commandPokedex(cfg *Config, words []string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.Pokemons {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
