package main

import (
	"fmt"
	"os"

	"github.com/senaphim/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*configuration) error
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

func commandExit(*configuration) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*configuration) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Println(fmt.Sprintf("%s: %s", cmd.name, cmd.description))
	}
	return nil
}

func commandMap(config *configuration) error {
	// Inserted this code to handle if you got to the end of the map but triggers on the first
	// TODO: Needs special casing on the first call ...
	// if config.nextUrl == nil {
	// 	fmt.Println("Congratulations. You've reached the last page!")
	// 	return nil
	// }

	locationsList, err := pokeapi.ListLocations(config.nextUrl)
	if err != nil {
		return err
	}

	config.nextUrl = locationsList.Next
	config.prevUrl = locationsList.Previous

	for _, location := range locationsList.Results {
		fmt.Println(fmt.Sprintf("%s", location.Name))
	}

	return nil
}

func commandMapb(config *configuration) error {
	if config.prevUrl == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locationsList, err := pokeapi.ListLocations(config.nextUrl)
	if err != nil {
		return err
	}

	config.nextUrl = locationsList.Next
	config.prevUrl = locationsList.Previous

	for _, location := range locationsList.Results {
		fmt.Println(fmt.Sprintf("%s", location.Name))
	}

	return nil
}
