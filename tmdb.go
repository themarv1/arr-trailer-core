package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TMDBVideoResult repräsentiert ein einzelnes Video aus der TMDB API-Antwort.
type TMDBVideoResult struct {
	Key      string `json:"key"`
	Site     string `json:"site"`
	Type     string `json:"type"`
	Official bool   `json:"official"`
	Language string `json:"iso_639_1"`
}

// TMDBVideosResponse ist die Gesamtstruktur der /videos API-Antwort von TMDB.
type TMDBVideosResponse struct {
	Results []TMDBVideoResult `json:"results"`
}

// findTrailerOnTMDB sucht nach einem passenden Trailer für einen Film auf TMDB.
// Gibt die YouTube-Video-ID oder einen leeren String zurück.
func findTrailerOnTMDB(movie Movie, apiKey string) (string, error) {
	// 1. Baue die API-URL zusammen
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d/videos?api_key=%s", movie.TmdbID, apiKey)

	// 2. Führe den HTTP GET Request aus
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request to TMDB failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("TMDB returned non-OK status: %s", resp.Status)
	}

	// 3. Parse die JSON-Antwort
	var videoResponse TMDBVideosResponse
	if err := json.NewDecoder(resp.Body).Decode(&videoResponse); err != nil {
		return "", fmt.Errorf("failed to parse TMDB JSON response: %w", err)
	}

	// 4. Finde den besten Trailer
	var bestTrailerKey string
	for _, video := range videoResponse.Results {
		// Wir wollen nur offizielle Trailer von YouTube
		if video.Site == "YouTube" && video.Type == "Trailer" {
			if video.Official {
				// Offizielle Trailer werden immer bevorzugt
				return video.Key, nil
			}
			// Speichere den ersten nicht-offiziellen Trailer als Fallback
			if bestTrailerKey == "" {
				bestTrailerKey = video.Key
			}
		}
	}

	return bestTrailerKey, nil
}
