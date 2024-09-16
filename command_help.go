package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
    fmt.Printf("\nWelcome to the Pokedex!\n")
    fmt.Printf("Usage:\n\n")

    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
    }

    fmt.Println()
    return nil
}
