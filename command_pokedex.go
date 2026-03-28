package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) != 1 {
		return errors.New("you have not caught any pokemon yet")
	}
	fmt.Println("Your Pokedex: ")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %v \n", pokemon.Name)
	}

	return nil
}
