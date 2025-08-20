package main

// DownloadConfig holds settings related to downloading trailers.
type DownloadConfig struct {
	Enabled   bool   `yaml:"enabled"`
	YTDLPPath string `yaml:"yt_dlp_path"`
	Quality   string `yaml:"quality"`
}

// Movie struct remains the same
type Movie struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Path      string `json:"path"`
	TmdbID    int    `json:"tmdbId"`
	HasFile   bool   `json:"hasFile"`
	Monitored bool   `json:"monitored"`
}
