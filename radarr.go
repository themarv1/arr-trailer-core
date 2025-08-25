package main

import (
	"fmt"
	"io"
	"net/http"
)

// getMovies performs an API call to a Radarr instance to fetch all movies.
// It returns the raw response body ([]byte) and a potential error.
func getMovies(instance RadarrInstance) ([]byte, error) {
	url := fmt.Sprintf("%s/api/v3/movie", instance.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("X-Api-Key", instance.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
