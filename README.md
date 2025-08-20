# Arr Trailer Core (ATC)

> **:warning: SUPER EARLY DEVELOPMENT ALPHA :warning:**
>
> This project is currently in a very early development stage. The core functionality of **detecting** missing trailers is working, but the actual **downloading** of trailers is not yet implemented. Please use it for testing and contribution purposes only. It is **not yet ready for production use**.

A powerful command-line tool designed to work with your Radarr (and soon Sonarr) library to ensure every movie has a locally stored trailer. ATC scans your library, identifies movies missing a trailer, and finds the correct trailer using The Movie Database (TMDB) for the highest accuracy.

## Features

-   **Radarr Library Scans:** Connects to one or more Radarr instances to process your entire movie library.
-   **Local Trailer Detection:** Reliably checks your movie folders for existing trailer files (e.g., `moviename-trailer.mkv`).
-   **High-Quality Trailer Search:** Uses the TMDB API to find official trailers for missing entries, ensuring high accuracy.
-   **Fallback Search:** If no trailer is found on TMDB, it can fall back to a direct search on YouTube.
-   **Flexible Path Mapping:** Intelligently translates paths between your Docker containers (like Radarr) and the host system where ATC is running.
-   **Highly Configurable:** Almost every aspect, from API keys to log levels, is controlled via a simple `config.yaml` file.
-   **Dry Run Mode:** Allows you to run the entire process in a simulation mode (`--dry-run`) to see what actions *would* be taken without making any changes.

## Configuration

To get started, you need to create a `config.yaml` file that contains your server details.

1.  Create your personal configuration file by copying the provided template:
    ```bash
    cp example-config.yaml config.yaml
    ```
2.  Open the new `config.yaml` file with a text editor.
3.  Fill in the actual URLs and API keys for your Radarr/Sonarr instances, your TMDB API key, and verify the path mappings. The `config.yaml` file is ignored by Git (`.gitignore`), so your private keys will remain safe.

## Usage

It is recommended to run ATC from a pre-compiled binary directly on your server (e.g., unRAID).

1.  [Build the binary for your server's operating system](#building-from-source).
2.  Copy the compiled `arr-trailer-core` binary and your `config.yaml` file to a directory on your server (e.g., `/mnt/user/appdata/arr-trailer-core/`).
3.  Make the binary executable (one-time command):
    ```bash
    chmod +x arr-trailer-core
    ```
4.  Run the program from the terminal:
    ```bash
    ./arr-trailer-core
    ```

### Command-Line Flags

-   `--config <path>`: Specify a custom path to your configuration file (default is `./config.yaml`).
-   `--dry-run`: Overrides the config file setting and forces a dry run.

## Building from Source

You can compile the project from source using the Go toolchain.

**To build from a linux operating system:**
```bash
go build -o arr-trailer-core .

**To cross-compile from Windows PowerShell:**
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o arr-trailer-core .

**To cross-compile from MacOS:**
GOOS=linux GOARCH=amd64 go build -o arr-trailer-core .

## Roadmap / Future Work

-   [ ] **Implement Trailer Downloader:** Integrate `yt-dlp` to perform the actual downloads.
-   [ ] **Full Sonarr Integration:** Implement the same detection and downloading logic for TV show episodes.
-   [ ] **Post-Processing:** Add optional steps to call `ffmpeg` for embedding metadata into the downloaded trailer files.
-   [ ] **Caching System:** Implement a local cache to avoid re-scanning and re-querying APIs for media that has been processed recently.
-   [ ] **Language Prioritization:** Allow users to define a preferred language list for trailer searches on TMDB.
-   [ ] **Interactive Setup:** A potential `setup` command to guide users through creating their first `config.yaml`.
