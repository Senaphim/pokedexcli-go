package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	// Declare map of allowed commands
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of locations in the Pokémon world. Calling the command agin will call the next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "As map, but returns to the previous list of locations",
			callback:    commandMapb,
		},
	}
	return commands
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Println(fmt.Sprintf("%s: %s", cmd.name, cmd.description))
	}
	return nil
}

func commandMap(*config) error {
	return nil
}

func commandMapb(*config) error {
	return nil
}
