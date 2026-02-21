package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal"
	"time"
)

// Initialize the cache with an interval of how long cache results are stored in time.Minute(s)
var locationCache = internal.NewCache(5 * time.Minute)

func fetchLocationData(url string) (pokemonLocationArea, error) {
	// Check Cache
	// Return Cache values if present
	// Make Fetch
	// Checks: - Did the call fail? - Did the Body read fail? - Is the status code bad? - Did the Unmarshal fail?
	// Close the body
	// Return unmarshaled struct

	data, ok := locationCache.Get(url)
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationData, fmt.Errorf("failed to read response body: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return locationData, fmt.Errorf("response failed with status code: %d, body: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return locationData, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return locationData, nil
}
