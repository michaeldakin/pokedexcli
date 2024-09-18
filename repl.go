package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/michaeldakin/pokedevcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		pokeRed := color.New(color.FgRed)
		pokeRed.Print("Pokedex > ")

		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading stdin:", err)
		}

		sanitised := sanitiseInput(scanner.Text())
		if len(sanitised) == 0 {
			continue
		}

		sanitisedArgLoc := sanitised
		// fmt.Printf("santised: %v\n", sanitised)

		cmdName := sanitised[0]
		cmd, ok := getCommands()[cmdName]
		if ok {
			err := cmd.Callback(sanitisedArgLoc, cfg)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func sanitiseInput(inputText string) []string {
	inputToLower := strings.ToLower(inputText)
	inputWords := strings.Fields(inputToLower)
	return inputWords
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func([]string, *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 locations",
			Callback:    commandMapForward,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the previous 20 locations",
			Callback:    commandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Find pokemon within the area",
			Callback:    exploreLocation,
		},
        "catch": {
            Name: "catch",
            Description: "Catch a pokemon and add it to your Pokedex",
            Callback: catch,
        },
		// "cache": {
		// 	Name:        "cache",
		// 	Description: "Print cache debug information",
		// 	Callback:    cacheDebug,
		// },
	}
}
