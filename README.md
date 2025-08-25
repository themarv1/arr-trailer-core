# Arr Trailer Core (ATC)

> **:warning: VERSION 1.0.0 :warning:**
>
> This is the first feature-complete version of the application. It can successfully detect, search for, and download missing trailers for Radarr and Sonarr. Please report any bugs or issues you encounter.

A powerful command-line tool designed to work with your Radarr and Sonarr libraries to ensure every movie and TV series has a locally stored trailer.

ATC scans your library, identifies media missing a trailer, finds the correct trailer using The Movie Database (TMDB), and downloads it directly into the correct folder using `yt-dlp` and `ffmpeg`.

## Features

-   **Radarr & Sonarr Support:** Connects to one or more Radarr and Sonarr instances to process your entire library.
-   **Local Trailer Detection:** Reliably checks your media folders for existing trailer files.
-   **High-Quality Trailer Search:** Uses the TMDB API to find official trailers for missing entries, ensuring high accuracy.
-   **Fallback Search:** If no trailer is found on TMDB, it falls back to a direct search on YouTube.
-   **Automated Downloading:** Integrates with `yt-dlp` to download the best-matched trailer directly into the correct movie or series folder.
-   **Flexible Path Mapping:** Intelligently translates paths between your Docker containers (like Radarr/Sonarr) and the host system where ATC is running.
-   **Highly Configurable:** Control every aspect via a simple `config.yaml` file, including API keys, log levels, download quality, and enabling/disabling features.
-   **Pre-flight Checks:** Verifies API connections, folder permissions, and dependencies (`yt-dlp`, `ffmpeg`) before starting to ensure a smooth run.
-   **Dry Run Mode:** Run the entire process in a simulation mode (`--dry-run`) to see what actions *would* be taken without downloading any files.

## Download & Installation

The easiest way to get started is by downloading the official release package for Linux (amd64).

1.  Go to the [**GitHub Releases**](https://github.com/themarv1/arr-trailer-core/releases/latest) page.
2.  Download the `arr-trailer-core-v1.0.0-linux-amd64.tar.gz` file.
3.  Upload the `.tar.gz` file to your server (e.g., to `/mnt/user/appdata/`).
4.  Unpack the archive using the terminal:
    ```bash
    tar -xzvf arr-trailer-core-v1.0.0-linux-amd64.tar.gz
    ```
    This will create a new folder named `arr-trailer-core-release` containing the application and all its dependencies.
5.  Navigate into the new directory:
    ```bash
    cd arr-trailer-core-release
    ```
6.  Follow the [Configuration](#configuration) steps below.

## Configuration

1.  Inside the `arr-trailer-core-release` folder, copy the example configuration file:
    ```bash
    cp example-config.yaml config.yaml
    ```
2.  Open `config.yaml` with a text editor (e.g., `nano config.yaml`) and fill in your details. The default paths for `yt-dlp` and `ffmpeg` are already set correctly to use the included files.

## Usage

1.  Make the main application executable (one-time command):
    ```bash
    chmod +x arr-trailer-core
    ```
2.  Run the program from within the `arr-trailer-core-release` directory:
    ```bash
    ./arr-trailer-core
    ```
    
### Command-Line Flags

-   `--config <path>`: Specify a custom path to your configuration file (default is `./config.yaml`).
-   `--dry-run`: Overrides the config file setting and forces a dry run.

## Building from Source

You can compile the project from source using the Go toolchain.

## Roadmap / Future Work

-   [ ] **Docker Container:** Create a `Dockerfile` for easy, self-contained deployment of the application and all its dependencies.
-   [ ] **Self-Updating Dependencies:** Implement a mechanism, especially for the Docker image, to automatically update `yt-dlp` to the latest version on startup.
-   [ ] **Post-Processing:** Add optional steps to call `ffmpeg` for embedding metadata into the downloaded trailer files.
-   [ ] **Caching System:** Implement a local cache to avoid re-scanning and re-querying APIs for media that has been processed recently.
-   [ ] **Language Prioritization:** Allow users to define a preferred language list for trailer searches on TMDB.
-   [ ] **Interactive Setup:** A potential `setup` command to guide users through creating their first `config.yaml`.
