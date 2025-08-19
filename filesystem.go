package main

import (
	"os"
	"strings"
)

// hasLocalTrailer prüft, ob im angegebenen Ordner eine Datei existiert,
// die auf einen Trailer hindeutet.
// Gibt true zurück, wenn ein Trailer gefunden wird, ansonsten false.
func hasLocalTrailer(movieFolderPath string) (bool, error) {
	// Lese den Inhalt des Verzeichnisses
	files, err := os.ReadDir(movieFolderPath)
	if err != nil {
		// Gibt einen Fehler zurück, wenn der Ordner nicht gelesen werden kann
		return false, err
	}

	// Gehe durch jede Datei im Ordner
	for _, file := range files {
		// Ignoriere Unterordner
		if file.IsDir() {
			continue
		}

		// Hole den Dateinamen
		fileName := file.Name()

		// Prüfe, ob der Dateiname unsere Trailer-Muster enthält
		// ToLower sorgt dafür, dass die Groß-/Kleinschreibung ignoriert wird
		if strings.Contains(strings.ToLower(fileName), "-trailer.") {
			return true, nil // Trailer gefunden!
		}
	}

	// Die Schleife ist durchgelaufen, kein Trailer wurde gefunden
	return false, nil
}
