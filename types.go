package main

// Movie repräsentiert die Struktur eines einzelnen Films aus der Radarr API-Antwort.
type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Path   string `json:"path"`
	TmdbID int    `json:"tmdbId"`
	// YouTubeTrailerID wurde entfernt
	HasFile   bool `json:"hasFile"`
	Monitored bool `json:"monitored"`
}
