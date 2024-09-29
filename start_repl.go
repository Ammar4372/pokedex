package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start_repl(cfg Config) {

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
			if len(words) == 2 {
				err := command.callback(&cfg, words[1])
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			err := command.callback(&cfg, "")
			if err != nil {
				fmt.Println(err)
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
