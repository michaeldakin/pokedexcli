package main

import (
	"os"
)

func commandExit(args []string, cfg *config) error {
    os.Exit(0)
    return nil
}
