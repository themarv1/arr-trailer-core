package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// sanitizeFilename removes characters that are invalid in file paths.
func sanitizeFilename(name string) string {
	replacer := strings.NewReplacer(
		"<", "", ">", "", ":", "-", "\"", "", "/", "\\", "|", "", "?", "", "*", "",
	)
	return replacer.Replace(name)
}

// downloadMovieTrailer handles the download for a specific movie.
func downloadMovieTrailer(movie Movie, youtubeKey string, config *Config, instance RadarrInstance) {
	var source string
	if youtubeKey != "" {
		source = "https://www.youtube.com/watch?v=" + youtubeKey
		log.Printf("[%s] INFO: Starting TMDB-found trailer download for '%s'", instance.Name, movie.Title)
	} else {
		source = fmt.Sprintf("ytsearch1:%s %d official trailer", movie.Title, movie.Year)
		log.Printf("[%s] INFO: Starting fallback trailer search/download for '%s'", instance.Name, movie.Title)
	}

	saneTitle := sanitizeFilename(movie.Title)
	fileNameTemplate := fmt.Sprintf("%s (%d)-trailer.%%(ext)s", saneTitle, movie.Year)
	translatedBasePath := translatePath(movie.Path, instance.PathMappings)
	outputPathTemplate := filepath.Join(translatedBasePath, fileNameTemplate)

	cmd := exec.Command(
		config.Download.YTDLPPath,
		"--ffmpeg-location", config.Download.FfmpegPath,
		"-f", config.Download.Quality,
		"-o", outputPathTemplate,
		"--no-mtime",
		source,
	)

	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] Running command: %s", cmd.String())
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("[%s] ERROR: yt-dlp failed for '%s'. Error: %v", instance.Name, movie.Title, err)
		log.Printf("[%s] yt-dlp output:\n%s", instance.Name, string(output))
		return
	}

	log.Printf("[%s] SUCCESS: Successfully downloaded trailer for '%s'.", instance.Name, movie.Title)
	if config.LogLevel == "debug" {
		log.Printf("[%s] yt-dlp output:\n%s", instance.Name, string(output))
	}
}

// downloadSeriesTrailer handles the download for a specific series.
func downloadSeriesTrailer(series Series, youtubeKey string, config *Config, instance SonarrInstance) {
	var source string
	if youtubeKey != "" {
		source = "https://www.youtube.com/watch?v=" + youtubeKey
		log.Printf("[%s] INFO: Starting TMDB-found trailer download for series '%s'", instance.Name, series.Title)
	} else {
		source = fmt.Sprintf("ytsearch1:%s official trailer", series.Title)
		log.Printf("[%s] INFO: Starting fallback trailer search/download for series '%s'", instance.Name, series.Title)
	}

	saneTitle := sanitizeFilename(series.Title)
	fileNameTemplate := fmt.Sprintf("%s-trailer.%%(ext)s", saneTitle)
	translatedBasePath := translatePath(series.Path, instance.PathMappings)
	trailersFolderPath := filepath.Join(translatedBasePath, "trailers")
	outputPathTemplate := filepath.Join(trailersFolderPath, fileNameTemplate)

	cmd := exec.Command(
		config.Download.YTDLPPath,
		"--ffmpeg-location", config.Download.FfmpegPath,
		"-f", config.Download.Quality,
		"-o", outputPathTemplate,
		"--no-mtime",
		source,
	)

	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] Running command: %s", cmd.String())
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("[%s] ERROR: yt-dlp failed for series '%s'. Error: %v", instance.Name, series.Title, err)
		log.Printf("[%s] yt-dlp output:\n%s", instance.Name, string(output))
		return
	}

	log.Printf("[%s] SUCCESS: Successfully downloaded trailer for series '%s'.", instance.Name, series.Title)
	if config.LogLevel == "debug" {
		log.Printf("[%s] yt-dlp output:\n%s", instance.Name, string(output))
	}
}
