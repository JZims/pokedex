package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	roster := cfg.caughtPokemon
	match, ok := roster[name]
	if !ok {
		return errors.New("you have not caught this pokemon")
	}
	fmt.Printf("Name: %v \n", match.Name)
	fmt.Printf("Height: %v \n", match.Height)
	fmt.Printf("Weight: %v \n", match.Weight)
	fmt.Println("Stats:")

	return nil
}
