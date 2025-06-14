package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commands := getCommands()

	// Infinite loop to read user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		commands[cleaned[0]].callback()
	}
}

func cleanInput(text string) []string {
	cleanString := strings.TrimSpace(text)
	cleanString = strings.ToLower(cleanString)
	cleanSlice := strings.Fields(cleanString)
	return cleanSlice
}
