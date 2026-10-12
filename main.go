package main

import "github.com/NamesNotRich/pokedexcli/internal/pokeapi"

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		commands: getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
