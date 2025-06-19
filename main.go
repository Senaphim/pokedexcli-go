package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/senaphim/pokedexcli/internal/pokecache"
)

type configuration struct {
	nextUrl *string
	prevUrl *string
}

func main() {
	var config configuration
	cache := pokecache.NewCache(5 * time.Second)
	commands := getCommands()

	// Infinite loop to read user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		var arg string
		if len(cleaned) > 1 {
			arg = cleaned[1]
		}
		err := commands[cleaned[0]].callback(&config, &cache, arg)
		if err != nil {
			fmt.Println(fmt.Errorf("Encountered error whilst running command %v: %v",
				commands[cleaned[0]].name, err))
		}
	}
}

func cleanInput(text string) []string {
	cleanString := strings.TrimSpace(text)
	cleanString = strings.ToLower(cleanString)
	cleanSlice := strings.Fields(cleanString)
	return cleanSlice
}
