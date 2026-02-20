package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal"
	"time"
)

type pokemonLocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Initialize the cache with an interval of how long cache results are stored in time.Minute(s)
var cache = internal.NewCache(5 * time.Minute)

func fetchData(url string) (pokemonLocationArea, error) {
	data, ok := cache.Get(url)
	if ok {
		var locationData pokemonLocationArea
		err := json.Unmarshal(data, &locationData)
		if err != nil {
			return locationData, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return locationData, nil
	}

	var locationData pokemonLocationArea
	res, err := http.Get(url)
	if err != nil {
		return locationData, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationData, fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode > 299 {
		return locationData, fmt.Errorf("response failed with status code: %d, body: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return locationData, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return locationData, nil
}
