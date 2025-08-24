package main

// DownloadConfig holds settings related to downloading trailers.
type DownloadConfig struct {
	Enabled    bool   `yaml:"enabled"`
	YTDLPPath  string `yaml:"yt_dlp_path"`
	FfmpegPath string `yaml:"ffmpeg_path"`
	Quality    string `yaml:"quality"`
}

// Movie represents the structure of a single movie from the Radarr API response.
type Movie struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Path      string `json:"path"`
	TmdbID    int    `json:"tmdbId"`
	HasFile   bool   `json:"hasFile"`
	Monitored bool   `json:"monitored"`
}

// Statistics represents the statistic data for a series or season.
type Statistics struct {
	EpisodeFileCount int `json:"episodeFileCount"`
}

// Season represents a season within a Sonarr series.
type Season struct {
	SeasonNumber int        `json:"seasonNumber"`
	Monitored    bool       `json:"monitored"`
	Statistics   Statistics `json:"statistics"`
}

// Series represents a single series from the Sonarr API response.
type Series struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Path       string     `json:"path"`
	TmdbID     int        `json:"tmdbId"`
	TvdbID     int        `json:"tvdbId"`
	Monitored  bool       `json:"monitored"`
	Seasons    []Season   `json:"seasons"`
	Statistics Statistics `json:"statistics"`
}
