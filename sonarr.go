package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// getSeries führt einen API-Aufruf an eine Sonarr-Instanz durch, um alle Serien abzurufen.
func getSeries(instance SonarrInstance) ([]Series, error) {
	url := fmt.Sprintf("%s/api/v3/series", instance.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for sonarr: %w", err)
	}

	req.Header.Set("X-Api-Key", instance.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to sonarr failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		statusErr := errors.New(resp.Status)
		// HIER IST DIE ÄNDERUNG: "Sonarr" -> "sonarr"
		return nil, fmt.Errorf("sonarr returned non-OK status: %w", statusErr)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read sonarr response body: %w", err)
	}

	var series []Series
	if err := json.Unmarshal(body, &series); err != nil {
		return nil, fmt.Errorf("failed to parse sonarr JSON response: %w", err)
	}

	return series, nil
}
