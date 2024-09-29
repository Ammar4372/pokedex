package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, words []string) error {
	if len(words) > 1 {
		fmt.Printf("Throwing a Pokeball at %s...\n", words[1])
		data, err := cfg.Client.GetPokemon(words[1])
		if err != nil {
			return err
		}
		num := rand.Intn(data.BaseExperience)
		if num > (data.BaseExperience)/2 {
			cfg.Pokemons[words[1]] = data
			fmt.Printf("%s was caught! \n", words[1])
			fmt.Println("You may now inspect it with the inspect command.")
			return nil
		}
		fmt.Printf("%s escaped \n", words[1])
		return nil
	}
	return fmt.Errorf("No Pokemon name was given")
}
