package main

import (
	"fmt"
)

func getPokedex(args []string, cfg *config) error {
	fmt.Printf("Pokedex contains %d pokemon\n", len(cfg.Pokedex))
	if len(cfg.Pokedex) >= 1 {
		for _, p := range cfg.Pokedex {
			fmt.Printf("- %s\n", p.Name)
		}
	}
	return nil
}
