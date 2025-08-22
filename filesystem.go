package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// hasLocalMovieTrailer checks if a file indicating a trailer exists in the given movie folder.
func hasLocalMovieTrailer(movieFolderPath string, config *Config) (bool, error) {
	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] Checking folder: %s", movieFolderPath)
	}
	files, err := os.ReadDir(movieFolderPath)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if config.LogLevel == "debug" {
			log.Printf("[DEBUG] ...seeing file: %s", file.Name())
		}
		lowerCaseFileName := strings.ToLower(file.Name())
		if strings.Contains(lowerCaseFileName, "trailer") {
			if config.LogLevel == "debug" {
				log.Printf("[DEBUG] HIT on file: %s", file.Name())
			}
			return true, nil
		}
	}
	return false, nil
}

// --- VERBESSERTE VERSION FÜR SONARR ---

// hasLocalSeriesTrailer checks for a 'trailers' subfolder (case-insensitive) and if it contains any files.
func hasLocalSeriesTrailer(seriesPath string, config *Config) (bool, error) {
	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] Checking for series trailer in base folder: %s", seriesPath)
	}

	// Read all entries in the series' root directory
	entries, err := os.ReadDir(seriesPath)
	if err != nil {
		return false, err
	}

	// Look for a directory that is likely to contain trailers (case-insensitive)
	var trailerDirName string
	for _, entry := range entries {
		if entry.IsDir() && strings.ToLower(entry.Name()) == "trailers" {
			trailerDirName = entry.Name() // Found it, store the actual name (e.g., "Trailers")
			break
		}
	}

	// If we didn't find a directory named 'trailers' (case-insensitive)
	if trailerDirName == "" {
		if config.LogLevel == "debug" {
			log.Printf("[DEBUG] ...no 'trailers' subfolder found.")
		}
		return false, nil
	}

	// We found the folder, now let's construct the full path to it
	trailerPath := filepath.Join(seriesPath, trailerDirName)
	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] ...found trailer subfolder at: %s", trailerPath)
	}

	// Now check if there are any files inside that folder
	trailerFiles, err := os.ReadDir(trailerPath)
	if err != nil {
		return false, err // Error reading the found trailers subfolder
	}

	for _, file := range trailerFiles {
		if !file.IsDir() {
			if config.LogLevel == "debug" {
				log.Printf("[DEBUG] ...found series trailer file: %s", file.Name())
			}
			return true, nil // Found at least one file, we're good
		}
	}

	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] ...'trailers' subfolder was empty.")
	}
	return false, nil // Folder was found, but it was empty
}
