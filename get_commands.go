package main

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
