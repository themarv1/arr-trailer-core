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

## Configuration

To get started, you need to create a `config.yaml` file that contains your server details.

1.  Create your personal configuration file by copying the provided template:
    ```bash
    cp example-config.yaml config.yaml
    ```
2.  Open the new `config.yaml` file with a text editor.
3.  Fill in the actual URLs and API keys for your Radarr/Sonarr instances, your TMDB API key, and verify all paths. The `config.yaml` file should be in your `.gitignore`, so your private keys will remain safe.

## Usage

It is recommended to run ATC from a pre-compiled binary directly on your server (e.g., unRAID).

1.  Download the `yt-dlp` executable for your system and place it in the same folder as `arr-trailer-core`.
2.  For best results, install `ffmpeg` on your system so `yt-dlp` can merge video and audio streams. On unRAID, this can be done via the "Nerd Pack" Community App.
3.  [Build the binary for your server's operating system](#building-from-source).
4.  Copy the compiled `arr-trailer-core` binary and your `config.yaml` file to a directory on your server.
5.  Make the binary executable (one-time command):
    ```bash
    chmod +x arr-trailer-core
    ```
6.  Run the program from the terminal:
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
```
**To cross-compile from Windows PowerShell:**
```Powershell
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o arr-trailer-core .
```
**To cross-compile from MacOS:**
```Bash
GOOS=linux GOARCH=amd64 go build -o arr-trailer-core .
```

## Roadmap / Future Work

-   [ ] **Docker Container:** Create a `Dockerfile` for easy, self-contained deployment of the application and all its dependencies.
-   [ ] **Self-Updating Dependencies:** Implement a mechanism, especially for the Docker image, to automatically update `yt-dlp` to the latest version on startup.
-   [ ] **Post-Processing:** Add optional steps to call `ffmpeg` for embedding metadata into the downloaded trailer files.
-   [ ] **Caching System:** Implement a local cache to avoid re-scanning and re-querying APIs for media that has been processed recently.
-   [ ] **Language Prioritization:** Allow users to define a preferred language list for trailer searches on TMDB.
-   [ ] **Interactive Setup:** A potential `setup` command to guide users through creating their first `config.yaml`.
