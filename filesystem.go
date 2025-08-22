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

// --- NEU FÜR SONARR ---

// hasLocalSeriesTrailer checks if a dedicated "trailers" subfolder with a video file exists.
func hasLocalSeriesTrailer(seriesPath string, config *Config) (bool, error) {
	trailerPath := filepath.Join(seriesPath, "trailers")

	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] Checking for series trailer folder: %s", trailerPath)
	}

	dirInfo, err := os.Stat(trailerPath)
	if os.IsNotExist(err) {
		return false, nil // Folder doesn't exist, so no trailer
	}
	if err != nil {
		return false, err // Another error occurred
	}
	if !dirInfo.IsDir() {
		return false, nil // It's a file, not a folder, so it's wrong
	}

	files, err := os.ReadDir(trailerPath)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if !file.IsDir() {
			if config.LogLevel == "debug" {
				log.Printf("[DEBUG] ...found series trailer file: %s", file.Name())
			}
			return true, nil
		}
	}
	return false, nil // Folder exists but is empty
}
