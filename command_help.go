package main

import (
	"fmt"
)

func commandHelp(cfg *Config, name string) error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the PokeDex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
