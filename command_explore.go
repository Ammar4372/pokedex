package main

import "fmt"

func commandExplore(cfg *Config, name string) error {
	data, err := cfg.Client.GetLocation(name)
	if err != nil {
		return err
	}
	for _, v := range data.Pokemons {
		fmt.Println(v.Pokemon.Name)
	}
	return nil
}
