package main

import (
	"strings"
)

func main() {
	println("Hello, World!")
}

func cleanInput(text string) []string {
	cleanString := strings.TrimSpace(text)
	cleanString = strings.ToLower(cleanString)
	cleanSlice := strings.Fields(cleanString)
	return cleanSlice
}
