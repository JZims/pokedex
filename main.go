package main

import (
	"bufio"
	"fmt"
	"os"
)

func getCommands(configData *config) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			config:      *configData,
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			config:      *configData,
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "A list of 20 location areas in the Pokemon world",
			config:      *configData,
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "If present, displays the previous 20 map results from the Pokemon world",
			config:      *configData,
			callback:    commandMapB,
		},
	}
}

func commandExit(configData *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(configData *config) error {
	fmt.Print("\nWelcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for name, description := range getCommands(&config{}) {
		fmt.Printf("%v: %v\n", name, description.description)
	}
	return nil
}

func commandMap(configData *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if configData.Next != nil && *configData.Next != "" {
		url = *configData.Next
	}

	locationData, err := fetchData(url)
	if err != nil {
		return err
	}

	// Update config with new Next/Previous from the flat structure
	if locationData.Next != "" {
		configData.Next = &locationData.Next
	} else {
		configData.Next = nil
	}
	if locationData.Previous != "" {
		configData.Previous = &locationData.Previous
	} else {
		configData.Previous = nil
	}

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(configData *config) error {
	if configData.Previous == nil || *configData.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locationData, err := fetchData(*configData.Previous)
	if err != nil {
		return err
	}

	// Update config with new Next/Previous from the flat structure
	if locationData.Next != "" {
		configData.Next = &locationData.Next
	} else {
		configData.Next = nil
	}
	if locationData.Previous != "" {
		configData.Previous = &locationData.Previous
	} else {
		configData.Previous = nil
	}

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	config      config
	callback    func(*config) error
}

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
		if len(input) == 0 {
			fmt.Print("No input given")
			continue
		}
		command, ok := commands[input[0]]
		if !ok {
			fmt.Print("Unknown command")
			continue
		}
		err := command.callback(&configData)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}

	}

}
