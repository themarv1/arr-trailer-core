package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// --- Structs for configuration ---
// (Diese sind jetzt alle in types.go, aber wir müssen die Config hier definieren)
type PathMapping struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}
type RadarrInstance struct {
	Name         string        `yaml:"name"`
	URL          string        `yaml:"url"`
	APIKey       string        `yaml:"api_key"`
	PathMappings []PathMapping `yaml:"path_mappings"`
}
type SonarrInstance struct {
	Name         string        `yaml:"name"`
	URL          string        `yaml:"url"`
	APIKey       string        `yaml:"api_key"`
	PathMappings []PathMapping `yaml:"path_mappings"`
}
type Config struct {
	LogLevel   string           `yaml:"log_level"`
	DryRun     bool             `yaml:"dry_run"`
	TmdbApiKey string           `yaml:"tmdb_api_key"`
	Download   DownloadConfig   `yaml:"download"`
	Radarr     []RadarrInstance `yaml:"radarr"`
	Sonarr     []SonarrInstance `yaml:"sonarr"`
}

// --- Helper function to translate paths ---
func translatePath(originalPath string, mappings []PathMapping) string {
	for _, mapping := range mappings {
		if strings.HasPrefix(originalPath, mapping.From) {
			return strings.Replace(originalPath, mapping.From, mapping.To, 1)
		}
	}
	return originalPath
}

func main() {
	// --- Setup ---
	log.Println("--- ATC v2.0 with instance-specific Path Mapping ---")
	configFile := flag.String("config", "config.yaml", "Path to the configuration file")
	cliDryRun := flag.Bool("dry-run", false, "Overrides the config file and forces a dry run.")
	flag.Parse()
	log.Println("Arr Trailer Core (ATC) is starting...")
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error loading configuration from '%s': %v", *configFile, err)
	}
	if config.LogLevel == "" {
		config.LogLevel = "info"
	}
	isDryRun := config.DryRun || *cliDryRun
	if isDryRun {
		log.Printf(">>> ATTENTION: Dry Run mode is active. No real changes will be made. <<<")
	}

	// --- Process Radarr Instances ---
	log.Println("Processing Radarr instances...")
	for _, instance := range config.Radarr {
		log.Printf("[%s] Fetching movie list...", instance.Name)
		movieData, err := getMovies(instance)
		if err != nil {
			log.Printf("[%s] ERROR fetching movies: %v", instance.Name, err)
			continue
		}
		var movies []Movie
		if err := json.Unmarshal(movieData, &movies); err != nil {
			log.Printf("[%s] ERROR parsing JSON response: %v", instance.Name, err)
			continue
		}
		log.Printf("[%s] Successfully found and processed %d movies.", instance.Name, len(movies))

		for _, movie := range movies {
			if !movie.Monitored || !movie.HasFile {
				continue
			}
			translatedPath := translatePath(movie.Path, instance.PathMappings)
			localTrailerFound, err := hasLocalMovieTrailer(translatedPath, config)
			if err != nil {
				log.Printf("[%s] WARNING: Could not check folder for '%s' ('%s'): %v", instance.Name, movie.Title, translatedPath, err)
				continue
			}
			if !localTrailerFound {
				log.Printf("[%s] MISSING: '%s (%d)' has no local trailer. Starting search...", instance.Name, movie.Title, movie.Year)
				youtubeKey := ""
				if config.TmdbApiKey != "" {
					key, err := findMovieTrailerOnTMDB(movie, config.TmdbApiKey)
					if err != nil {
						log.Printf("[%s] INFO: TMDB search for '%s' failed: %v. Falling back to direct search.", instance.Name, movie.Title, err)
					}
					youtubeKey = key
				}
				if youtubeKey != "" {
					log.Printf("[%s] ACTION: Found trailer for '%s' on TMDB (YouTube Key: %s). Would download now.", instance.Name, movie.Title, youtubeKey)
				} else {
					log.Printf("[%s] ACTION: No trailer found on TMDB for '%s'. Would use direct yt-dlp search now.", instance.Name, movie.Title)
				}
			} else {
				log.Printf("[%s] OK: '%s (%d)' already has a local trailer.", instance.Name, movie.Title, movie.Year)
			}
		}
	}

	// --- Process Sonarr Instances ---
	log.Println("Processing Sonarr instances...")
	for _, instance := range config.Sonarr {
		log.Printf("[%s] Fetching series list...", instance.Name)
		seriesList, err := getSeries(instance)
		if err != nil {
			log.Printf("[%s] ERROR fetching series: %v", instance.Name, err)
			continue
		}
		log.Printf("[%s] Successfully found %d series.", instance.Name, len(seriesList))

		for _, series := range seriesList {
			if !series.Monitored {
				continue
			}
			translatedPath := translatePath(series.Path, instance.PathMappings)
			localTrailerFound, err := hasLocalSeriesTrailer(translatedPath, config)
			if err != nil {
				log.Printf("[%s] WARNING: Could not check series folder for '%s' ('%s'): %v", instance.Name, series.Title, translatedPath, err)
				continue
			}
			if !localTrailerFound {
				log.Printf("[%s] MISSING: Series '%s' has no local trailer. Starting search...", instance.Name, series.Title)
				youtubeKey := ""
				if config.TmdbApiKey != "" {
					key, err := findSeriesTrailerOnTMDB(series, config.TmdbApiKey)
					if err != nil {
						log.Printf("[%s] INFO: TMDB search for '%s' failed: %v. Falling back to direct search.", instance.Name, series.Title, err)
					}
					youtubeKey = key
				}
				if youtubeKey != "" {
					log.Printf("[%s] ACTION: Found trailer for '%s' on TMDB (YouTube Key: %s). Would download now.", instance.Name, series.Title, youtubeKey)
				} else {
					log.Printf("[%s] ACTION: No trailer found on TMDB for '%s'. Would use direct yt-dlp search now.", instance.Name, series.Title)
				}
			} else {
				log.Printf("[%s] OK: Series '%s' already has a local trailer.", instance.Name, series.Title)
			}
		}
	}

	log.Println("Arr Trailer Core (ATC) has finished.")
}

// --- Helper function to load the YAML file ---
func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
