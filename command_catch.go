package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(args []string, cfg *config) error {
	// fmt.Printf("DEBUG exploreLocation: %v\n", argLoc)

	if len(args) < 2 {
		return errors.New("no pokemon provided")
	}
	if len(args) > 2 {
		return errors.New("only 1 pokemon required")
	}

	pokemonName := args[1]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonStats, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	baseChanceLimit := int(float64(pokemonStats.BaseExperience) * 1.25)
	catchPokemonRand := rand.Intn(baseChanceLimit)

	if catchPokemonRand >= pokemonStats.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.Pokedex[pokemonStats.Name] = pokemonStats
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
}
