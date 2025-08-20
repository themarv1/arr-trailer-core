package main

import (
	"log"
	"os"
	"strings"
)

// hasLocalTrailer prüft, ob im angegebenen Ordner eine Datei existiert,
// die auf einen Trailer hindeutet.
// Die Funktion akzeptiert jetzt die Config, um das Loglevel zu prüfen.
func hasLocalTrailer(movieFolderPath string, config *Config) (bool, error) {
	if config.LogLevel == "debug" {
		log.Printf("[DEBUG] Prüfe Ordner: %s", movieFolderPath)
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
			log.Printf("[DEBUG] ...sehe Datei: %s", file.Name())
		}

		lowerCaseFileName := strings.ToLower(file.Name())

		if strings.Contains(lowerCaseFileName, "trailer") {
			if config.LogLevel == "debug" {
				log.Printf("[DEBUG] TREFFER bei Datei: %s", file.Name())
			}
			return true, nil
		}
	}

	return false, nil
}
