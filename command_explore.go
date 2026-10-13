package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("you have to provide exactly one city name")
	}

	city := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(city)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", city)
	fmt.Println("Found Pokemon:")
	for _, pEncounter := range resp.PokemonEncounters {
		fmt.Printf("- %v\n", pEncounter.Pokemon.Name)
	}
	return nil
}
