package main

import (
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
		"explore": {
			name:        "explore",
			description: "Displays a list of available Pokemon in the area given as a parameter",
			config:      *configData,
			callback:    commandExplore,
		},
	}
}

func commandExit(configData *config, areaName string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(configData *config, areaName string) error {
	fmt.Print("\nWelcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for name, description := range getCommands(&config{}) {
		fmt.Printf("%v: %v\n", name, description.description)
	}
	return nil
}

func commandMap(configData *config, areaName string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if configData.Next != nil && *configData.Next != "" {
		url = *configData.Next
	}

	locationData, err := fetchLocationData(url)
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

func commandMapB(configData *config, areaName string) error {
	if configData.Previous == nil || *configData.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locationData, err := fetchLocationData(*configData.Previous)
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

func commandExplore(configData *config, areaName string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", areaName)
	pokemonData, err := fetchPokemonData(url)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonData.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
