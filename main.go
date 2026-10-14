package main

import (
	"time"

	"github.com/NamesNotRich/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Second)
	cfg := &config{
		commands:      getCommands(),
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
