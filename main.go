package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		fmt.Println(fmt.Sprintf("Your command was: %v", cleaned[0]))
	}
}

func cleanInput(text string) []string {
	cleanString := strings.TrimSpace(text)
	cleanString = strings.ToLower(cleanString)
	cleanSlice := strings.Fields(cleanString)
	return cleanSlice
}
