package main


import (
	"errors"
	"fmt"
)

func exploreLocation(argLoc []string, cfg *config) error {
	// fmt.Printf("DEBUG exploreLocation: %v\n", argLoc)

	if argLoc[0] == "explore" {
		if len(argLoc) < 2 {
			return errors.New("no location provided")
		}
		if len(argLoc) > 2 {
			return errors.New("only 1 location required")
		}
	}

    exploreAreaLoc := argLoc[1]

    fmt.Printf("Exploring %s...\n", exploreAreaLoc)

	locPokemon, err := cfg.pokeapiClient.GetLocationPokemon(exploreAreaLoc)
	if err != nil {
		return err
	}

    if len(locPokemon.PokemonEncounters) < 1 {
        fmt.Printf("No pokemon found at %s\n", exploreAreaLoc)
        return nil
    }
	for _, encounter := range locPokemon.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}

	return nil
}
