package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start_repl() {
	for {
		fmt.Printf("PokeDex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cmd := words[0]
		command, exists := getCommands()[cmd]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println("err: %w", err)
			}
			continue

		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanInput(cmd string) []string {
	output := strings.ToLower(cmd)
	words := strings.Fields(output)
	return words
}
