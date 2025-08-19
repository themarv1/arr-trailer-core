package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// --- Structs für die Konfiguration ---
type RadarrInstance struct {
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
	APIKey string `yaml:"api_key"`
}

type SonarrInstance struct {
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
	APIKey string `yaml:"api_key"`
}

type Config struct {
	DryRun bool             `yaml:"dry_run"`
	Radarr []RadarrInstance `yaml:"radarr"`
	Sonarr []SonarrInstance `yaml:"sonarr"`
}

func main() {
	// --- Kommandozeilen-Flags ---
	configFile := flag.String("config", "config.yaml", "Pfad zur Konfigurationsdatei")
	cliDryRun := flag.Bool("dry-run", false, "Überschreibt die Konfiguration und erzwingt einen Dry Run.")
	flag.Parse()

	log.Println("Arr Trailer Core (ATC) startet...")

	// --- Konfiguration laden ---
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Fehler beim Laden der Konfiguration von '%s': %v", *configFile, err)
	}

	// --- Dry-Run-Status bestimmen und loggen ---
	isDryRun := config.DryRun || *cliDryRun
	if isDryRun {
		reason := "in der Konfigurationsdatei aktiviert"
		if *cliDryRun {
			reason = "durch --dry-run Flag erzwungen"
		}
		log.Printf(">>> ACHTUNG: Dry Run Modus ist aktiviert (%s). Es werden keine echten Änderungen vorgenommen. <<<", reason)
	}

	// --- Radarr-Instanzen verarbeiten ---
	log.Println("Verarbeite Radarr-Instanzen...")
	for _, instance := range config.Radarr {
		log.Printf("[%s] Rufe Filmliste ab...", instance.Name)
		movieData, err := getMovies(instance)
		if err != nil {
			log.Printf("[%s] FEHLER beim Abrufen der Filme: %v", instance.Name, err)
			continue // Mache mit der nächsten Instanz weiter
		}

		var movies []Movie
		if err := json.Unmarshal(movieData, &movies); err != nil {
			log.Printf("[%s] FEHLER beim Verarbeiten der JSON-Antwort: %v", instance.Name, err)
			continue
		}
		log.Printf("[%s] Erfolgreich %d Filme gefunden und verarbeitet.", instance.Name, len(movies))

		// --- VEREINFACHTE PRÜF-LOGIK ---
		for _, movie := range movies {
			// Wir interessieren uns nur für überwachte Filme
			if !movie.Monitored {
				continue
			}

			// Prüfe den lokalen Ordner des Films auf eine Trailer-Datei
			localTrailerFound, err := hasLocalTrailer(movie.Path)
			if err != nil {
				log.Printf("[%s] WARNUNG: Konnte den Ordner für '%s' nicht prüfen: %v", instance.Name, movie.Title, err)
				continue
			}

			// Wenn KEIN lokaler Trailer gefunden wurde, logge die Meldung.
			if !localTrailerFound {
				log.Printf("[%s] Film '%s (%d)' hat keinen lokalen Trailer.", instance.Name, movie.Title, movie.Year)
			}
		}

		// Hier kommt später die Logik für Aktionen im Live-Modus hin
		if isDryRun {
			log.Printf("[%s] DRY RUN: Es würden jetzt Aktionen für die Filme ausgeführt.", instance.Name)
		} else {
			// Echte Aktionen...
		}
	}

	// --- Sonarr-Instanzen verarbeiten (Platzhalter) ---
	log.Println("Verarbeite Sonarr-Instanzen...")
	for _, instance := range config.Sonarr {
		log.Printf(" - Verarbeite Instanz: %s (%s)", instance.Name, instance.URL)
	}

	log.Println("Arr Trailer Core (ATC) hat den Vorgang abgeschlossen.")
}

// --- Hilfsfunktion zum Laden der YAML-Datei ---
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
