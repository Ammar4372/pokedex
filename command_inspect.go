package main

import "fmt"

func commandInspect(cfg *Config, words []string) error {
	if len(words) > 1 {
		if pokemon, ok := cfg.Pokemons[words[1]]; ok {
			fmt.Printf("Name : %s\n", pokemon.Name)
			fmt.Printf("Heignt : %d\n", pokemon.Height)
			fmt.Printf("Weight : %d\n", pokemon.Weight)
			fmt.Println("Stats:")
			for _, stat := range pokemon.Stats {
				fmt.Printf(" - %s: %d \n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types:")
			for _, v := range pokemon.Types {
				fmt.Printf(" - %s\n", v.Type.Name)
			}
			return nil
		}
		return fmt.Errorf("You haven't caught %s yet", words[1])
	}
	return fmt.Errorf("No Pokemon name was given")
}
