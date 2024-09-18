package main

import (
	"errors"
	"fmt"
)

func inspect(args []string, cfg *config) error {
	if len(args) < 2 {
		return errors.New("no pokemon provided")
	}
	if len(args) > 2 {
		return errors.New("only 1 pokemon required")
	}

	inspectPokemonName := args[1]
	val, ok := cfg.Pokedex[inspectPokemonName]

	if ok {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Stats:\n")
        for _, v := range val.Stats {
            fmt.Printf("  - %s: %d\n", v.Stat.Name, v.BaseStat)
        }
        fmt.Printf("Types:\n")
        for _, v := range val.Types {
            fmt.Printf("  - %s\n", v.Type.Name)
        }
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
}
