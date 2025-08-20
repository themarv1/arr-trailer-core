package main

import (
	"os"
	"path/filepath" // Neu für die Dateiendung
	"strings"
)

// hasLocalTrailer prüft, ob im angegebenen Ordner eine Datei existiert,
// die auf einen Trailer hindeutet.
// Gibt true zurück, wenn ein Trailer gefunden wird, ansonsten false.
func hasLocalTrailer(movieFolderPath string) (bool, error) {
	files, err := os.ReadDir(movieFolderPath)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		// --- NEUE, ROBUSTERE PRÜFUNG ---
		// 1. Dateiendung entfernen
		extension := filepath.Ext(fileName)
		fileNameWithoutExt := strings.TrimSuffix(fileName, extension)

		// 2. Umwandeln in Kleinbuchstaben und Leerzeichen am Ende entfernen
		normalizedFileName := strings.TrimSpace(strings.ToLower(fileNameWithoutExt))

		// 3. Prüfen, ob der normalisierte Name auf "-trailer" endet
		if strings.HasSuffix(normalizedFileName, "-trailer") {
			return true, nil // Trailer gefunden!
		}
	}

	return false, nil
}
