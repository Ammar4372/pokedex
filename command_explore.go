package main

import "fmt"

func commandExplore(cfg *Config, name string) error {
	fmt.Printf("Exploring %s...\n", name)
	data, err := cfg.Client.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemons:")
	for _, v := range data.Pokemons {
		fmt.Printf(" - %s \n", v.Pokemon.Name)
	}
	return nil
}
