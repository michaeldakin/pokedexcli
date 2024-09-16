package main

import (
	"errors"
	"fmt"
)

func commandMapForward(cfg *config) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locations.Next
	cfg.prevLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.prevLocation == nil {
		return errors.New("unable to go back, on first page")
	}

	locations, err := cfg.pokeapiClient.GetLocations(cfg.prevLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locations.Next
	cfg.prevLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
