package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading stdin:", err)
		}

        sanitised := sanitiseInput(scanner.Text())
        if len(sanitised) == 0 {
            continue
        }

        cmdName := sanitised[0]
        cmd, ok := getCommands()[cmdName]
        if ok {
            err := cmd.callback()
            if err != nil {
                fmt.Println(err)
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
