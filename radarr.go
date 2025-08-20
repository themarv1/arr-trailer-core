package main

import (
	"fmt"
	"io"
	"net/http"
)

// getMovies führt einen API-Aufruf an eine Radarr-Instanz durch, um alle Filme abzurufen.
// Die Funktion gibt die rohen Antwortdaten ([]byte) und einen eventuellen Fehler zurück.
func getMovies(instance RadarrInstance) ([]byte, error) {
	url := fmt.Sprintf("%s/api/v3/movie", instance.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Erstellen des Requests: %w", err)
	}

	req.Header.Set("X-Api-Key", instance.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Senden des Requests: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unerwarteter Status Code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Lesen der Antwort: %w", err)
	}

	return body, nil
}
