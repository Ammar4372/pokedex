package main

import "fmt"

func commandExplore(cfg *Config, words []string) error {
	if len(words) > 1 {
		fmt.Printf("Exploring %s...\n", words[1])
		data, err := cfg.Client.GetLocation(words[1])
		if err != nil {
			return err
		}
		fmt.Println("Found Pokemons:")
		for _, v := range data.PokemonEncounters {
			fmt.Printf(" - %s \n", v.Pokemon.Name)
		}
		return nil
	}
	return fmt.Errorf("No location name was given")
}
