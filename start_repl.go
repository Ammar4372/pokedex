package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start_repl(cfg Config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		cmd := words[0]
		command, exists := getCommands()[cmd]
		if exists {
			err := command.callback(&cfg, words)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(cmd string) []string {
	output := strings.ToLower(cmd)
	words := strings.Fields(output)
	return words
}
