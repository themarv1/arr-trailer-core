package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// RootFolder represents a single root folder from the Radarr/Sonarr API
type RootFolder struct {
	Path string `json:"path"`
}

// checkArrConnection checks the API connection to a Radarr or Sonarr instance.
func checkArrConnection(url, apiKey, arrType string) error {
	apiPath := "/api/v3/system/status"
	fullURL := fmt.Sprintf("%s%s", url, apiPath)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("returned non-OK status: %s", resp.Status)
	}

	log.Printf("INFO: Successfully connected to %s instance at %s", arrType, url)
	return nil
}

// getRootFolders fetches the list of root folders from a Radarr or Sonarr instance.
func getRootFolders(url, apiKey, arrType string) ([]string, error) {
	apiPath := "/api/v3/rootfolder"
	fullURL := fmt.Sprintf("%s%s", url, apiPath)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for root folders: %w", err)
	}
	req.Header.Set("X-Api-Key", apiKey)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request for root folders failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("returned non-OK status for root folders: %s", resp.Status)
	}
	var rootFolders []RootFolder
	if err := json.NewDecoder(resp.Body).Decode(&rootFolders); err != nil {
		return nil, fmt.Errorf("failed to parse root folders JSON: %w", err)
	}

	var paths []string
	for _, folder := range rootFolders {
		paths = append(paths, folder.Path)
	}
	return paths, nil
}

// checkWritePermissions checks write access for a given list of paths.
func checkWritePermissions(paths []string) error {
	uniquePaths := make(map[string]bool)
	for _, path := range paths {
		uniquePaths[path] = true
	}

	for path := range uniquePaths {
		log.Printf("INFO: Checking write permissions for path '%s'...", path)
		testFilePath := filepath.Join(path, ".atc-writetest")

		file, err := os.Create(testFilePath)
		if err != nil {
			return fmt.Errorf("failed to create test file in '%s': %w. Please check permissions", path, err)
		}
		file.Close()

		err = os.Remove(testFilePath)
		if err != nil {
			return fmt.Errorf("failed to remove test file in '%s': %w. Please check permissions", path, err)
		}
	}
	return nil
}

// checkDependencies verifies that required external command-line tools are installed.
func checkDependencies(config *DownloadConfig) error {
	// Check for yt-dlp
	log.Printf("INFO: Checking for dependency 'yt-dlp' using path '%s'...", config.YTDLPPath)
	_, err := exec.LookPath(config.YTDLPPath)
	if err != nil {
		return fmt.Errorf("yt-dlp executable not found at path '%s'. Please ensure it is installed and the path in your config.yaml is correct: %w", config.YTDLPPath, err)
	}
	log.Printf("INFO: Dependency 'yt-dlp' found successfully.")

	// Check for ffmpeg
	log.Printf("INFO: Checking for dependency 'ffmpeg' using path '%s'...", config.FfmpegPath)
	_, err = exec.LookPath(config.FfmpegPath)
	if err != nil {
		return fmt.Errorf("ffmpeg executable not found at path '%s'. yt-dlp requires ffmpeg to merge video and audio for the best quality. Please install ffmpeg on your system (e.g., via Nerd Pack on unRAID) or download a static build from https://ffmpeg.org/download.html and provide the correct path in your config.yaml: %w", config.FfmpegPath, err)
	}
	log.Printf("INFO: Dependency 'ffmpeg' found successfully.")

	return nil
}
