package main

func getCommands() map[string]command {
	return map[string]command{
		"explore": {
			name:        "explore",
			description: "explore a area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all Pokemon you have caught",
			callback:    commandPokedex,
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
