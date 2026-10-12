package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	resp, err := cfg.pokeapiClient.FetchLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.FetchLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
