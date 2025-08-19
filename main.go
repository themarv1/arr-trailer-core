package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config repräsentiert die Konfigurationsstruktur der YAML-Datei
type Config struct {
	Radarr struct {
		URL    string `yaml:"url"`
		APIKey string `yaml:"api_key"`
	} `yaml:"radarr"`
	Sonarr struct {
		URL    string `yaml:"url"`
		APIKey string `yaml:"api_key"`
	} `yaml:"sonarr"`
}

func main() {
	log.Println("Arr Trailer Core (ATC) startet...")

	config, err := loadConfig("config.example.yaml")
	if err != nil {
		log.Fatalf("Fehler beim Laden der Konfiguration: %v", err)
	}

	log.Printf("Radarr-URL aus Konfiguration: %s", config.Radarr.URL)
	log.Printf("Sonarr-URL aus Konfiguration: %s", config.Sonarr.URL)

	log.Println("Arr Trailer Core (ATC) hat den Vorgang abgeschlossen.")
}

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
