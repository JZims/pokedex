package main

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	config      config
	callback    func(*config, ...string) error
}

type pokemonLocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokemonEncounter struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
