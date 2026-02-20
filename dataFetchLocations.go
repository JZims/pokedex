package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func fetchData(url string) (pokemonLocationArea, error) {
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
