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
		args := []string{}

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
		if input[0] == "explore" && len(input) < 2 {
			fmt.Println("No area argument provided. Try again.")
			continue
		}
		args = append(args, input...)

		err := command.callback(&configData, args...)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}

	}

}
