package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TMDBVideoResult represents a single video from the TMDB API response.
type TMDBVideoResult struct {
	Key      string `json:"key"`
	Site     string `json:"site"`
	Type     string `json:"type"`
	Official bool   `json:"official"`
	Language string `json:"iso_639_1"`
}

// TMDBVideosResponse is the top-level structure of the /videos API response from TMDB.
type TMDBVideosResponse struct {
	Results []TMDBVideoResult `json:"results"`
}

// findMovieTrailerOnTMDB searches for a suitable trailer for a movie on TMDB.
func findMovieTrailerOnTMDB(movie Movie, apiKey string) (string, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d/videos?api_key=%s", movie.TmdbID, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request to TMDB failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("TMDB returned non-OK status: %s", resp.Status)
	}
	var videoResponse TMDBVideosResponse
	if err := json.NewDecoder(resp.Body).Decode(&videoResponse); err != nil {
		return "", fmt.Errorf("failed to parse TMDB JSON response: %w", err)
	}

	var bestTrailerKey string
	for _, video := range videoResponse.Results {
		if video.Site == "YouTube" && video.Type == "Trailer" {
			if video.Official {
				return video.Key, nil
			}
			if bestTrailerKey == "" {
				bestTrailerKey = video.Key
			}
		}
	}
	return bestTrailerKey, nil
}

// --- NEU FÜR SONARR ---

// findSeriesTrailerOnTMDB searches for a suitable trailer for a series on TMDB.
func findSeriesTrailerOnTMDB(series Series, apiKey string) (string, error) {
	// The API endpoint for TV shows is /tv/{id}/videos
	url := fmt.Sprintf("https://api.themoviedb.org/3/tv/%d/videos?api_key=%s", series.TmdbID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request to TMDB failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("TMDB returned non-OK status: %s", resp.Status)
	}

	var videoResponse TMDBVideosResponse
	if err := json.NewDecoder(resp.Body).Decode(&videoResponse); err != nil {
		return "", fmt.Errorf("failed to parse TMDB JSON response: %w", err)
	}

	var bestTrailerKey string
	for _, video := range videoResponse.Results {
		if video.Site == "YouTube" && video.Type == "Trailer" {
			if video.Official {
				return video.Key, nil
			}
			if bestTrailerKey == "" {
				bestTrailerKey = video.Key
			}
		}
	}
	return bestTrailerKey, nil
}
