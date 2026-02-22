package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal"
	"time"
)

var exploreCache = internal.NewCache(5 * time.Minute)

func fetchPokemonData(url string) (pokemonEncounter, error) {

	data, ok := exploreCache.Get(url)
	if ok {
		var exploreData pokemonEncounter
		err := json.Unmarshal(data, &exploreData)
		if err != nil {
			return exploreData, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return exploreData, nil
	}

	var exploreData pokemonEncounter
	res, err := http.Get(url)
	if err != nil {
		return exploreData, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return exploreData, fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode > 299 {
		return exploreData, fmt.Errorf("response failed with status code: %d, body: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &exploreData)
	if err != nil {
		return exploreData, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return exploreData, nil
}
