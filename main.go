package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	configData := config{
		Next:     nil,
		Previous: nil,
	}
	commands := getCommands(&configData)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)
		arg := ""

		if len(input) == 0 {
			fmt.Print("No input given")
			continue
		}
		command, ok := commands[input[0]]
		if !ok {
			fmt.Print("Unknown command")
			continue
		}

		// Explore Command-specific
		if input[0] == "explore" && len(input) > 1 {
			location := input[1]
			if location != "" {
				arg = input[1]
			} else {
				fmt.Printf("'%v' is not a valid location. Try again.", input[1])
				continue
			}

		}

		err := command.callback(&configData, arg)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}

	}

}
