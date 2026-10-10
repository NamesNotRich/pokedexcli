package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command := input[0]

		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	strSlice := strings.Fields(loweredText)
	return strSlice
}
