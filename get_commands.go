package main

func getCommands() map[string]command {
	return map[string]command{
		"explore": {
			name:        "explore",
			description: "explore a area",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Display next twenty location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous twenty location",
			callback:    commandMapb,
		},

		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
