package main

// Movie repräsentiert die Struktur eines einzelnen Films aus der Radarr API-Antwort.
// Die `json:"..."`-Tags sagen Go, wie die Felder in der JSON-Datei heißen.
type Movie struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Path      string `json:"path"`
	TmdbID    int    `json:"tmdbId"`
	HasFile   bool   `json:"hasFile"`
	Monitored bool   `json:"monitored"`
}
