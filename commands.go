package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/senaphim/pokedexcli/internal/pokeapi"
	"github.com/senaphim/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*configuration, *pokecache.Cache, []string) error
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
		"explore": {
			name:        "explore",
			description: "Takes one argument of location. Displays pokeman catchable at that location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon. Chance to catch based off of level of pokemon. Takes one pokemon as an argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays information about caught pokemon. Takes one pokemon as an argument",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all caught pokemon",
			callback:    commandPokedex,
		},
	}
	return commands
}

func commandExit(*configuration, *pokecache.Cache, []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*configuration, *pokecache.Cache, []string) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Println(fmt.Sprintf("%s: %s", cmd.name, cmd.description))
	}
	return nil
}

func commandMap(config *configuration, cache *pokecache.Cache, _ []string) error {
	// Inserted this code to handle if you got to the end of the map but triggers on the first
	// TODO: Needs special casing on the first call ...
	// if config.nextUrl == nil {
	// 	fmt.Println("Congratulations. You've reached the last page!")
	// 	return nil
	// }

	locationsList, err := pokeapi.ListLocations(config.nextUrl, cache)
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

func commandMapb(config *configuration, cache *pokecache.Cache, _ []string) error {
	if config.prevUrl == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locationsList, err := pokeapi.ListLocations(config.nextUrl, cache)
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

func commandExplore(_ *configuration, cache *pokecache.Cache, location []string) error {
	locationDetails, err := pokeapi.ExploreLocation(location[0], cache)
	if err != nil {
		return err
	}

	for _, mon := range locationDetails.PokemonEncounters {
		fmt.Println(fmt.Sprintf("%s", mon.Pokemon.Name))
	}

	return nil
}

func commandCatch(conf *configuration, cache *pokecache.Cache, mon []string) error {
	pokemon, err := pokeapi.GetPokemon(mon[0], cache)
	if err != nil {
		return err
	}

	if mon[1] == "masterball" {
		fmt.Println(fmt.Sprintf("%v was caught!", pokemon.Species.Name))
		conf.caught[pokemon.Species.Name] = pokemon
		return nil
	}

	fmt.Println(fmt.Sprintf("Throwing a Pokeball at %s...", pokemon.Species.Name))

	catchInt := rand.IntN(100)
	if catchInt > pokemon.BaseExperience {
		fmt.Println(fmt.Sprintf("%v was caught!", pokemon.Species.Name))
		conf.caught[pokemon.Species.Name] = pokemon
	} else {
		fmt.Println(fmt.Sprintf("%v escaped!", pokemon.Species.Name))
	}

	return nil
}

func commandInspect(conf *configuration, _ *pokecache.Cache, mon []string) error {
	pokemon, ok := conf.caught[mon[0]]
	if !ok {
		fmt.Println("You have not caught that pokemon yet")
		return nil
	}

	fmt.Println(fmt.Sprintf("Name: %s", pokemon.Species.Name))
	fmt.Println(fmt.Sprintf("Height: %v", pokemon.Height))
	fmt.Println(fmt.Sprintf("Weight: %v", pokemon.Height))
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Println(fmt.Sprintf("  -%v: %v", stat.Stat.Name, stat.BaseStat))
	}
	fmt.Println("Types:")
	for _, monType := range pokemon.Types {
		fmt.Println(fmt.Sprintf("  -%v", monType.Type.Name))
	}
	return nil
}

func commandPokedex(conf *configuration, _ *pokecache.Cache, _ []string) error {
	fmt.Println("Your pokedex:")
	for mon := range conf.caught {
		fmt.Println(fmt.Sprintf("  -%s", mon))
	}

	return nil
}
