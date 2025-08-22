# Arr Trailer Core (ATC)

A powerful command-line tool designed to work with your Radarr and Sonarr libraries to ensure every movie and TV series has a locally stored trailer.

ATC scans your library, identifies media missing a trailer, finds the correct trailer using The Movie Database (TMDB), and downloads it directly into the correct folder using `yt-dlp`.

## Features

-   **Radarr & Sonarr Support:** Connects to one or more Radarr and Sonarr instances to process your entire library.
-   **Local Trailer Detection:** Reliably checks your media folders for existing trailer files.
-   **High-Quality Trailer Search:** Uses the TMDB API to find official trailers for missing entries, ensuring high accuracy.
-   **Fallback Search:** If no trailer is found on TMDB, it falls back to a direct search on YouTube.
-   **Automated Downloading:** Integrates with `yt-dlp` to download the best-matched trailer directly into the correct movie or series folder.
-   **Flexible Path Mapping:** Intelligently translates paths between your Docker containers (like Radarr/Sonarr) and the host system where ATC is running.
-   **Highly Configurable:** Control every aspect via a simple `config.yaml` file, including API keys, log levels, download quality, and enabling/disabling features.
-   **Dry Run Mode:** Run the entire process in a simulation mode (`--dry-run`) to see what actions *would* be taken without downloading any files.

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

1.  Ensure `yt-dlp` is installed on your server and accessible via the system's PATH or provide a direct path in the `config.yaml`.
2.  [Build the binary for your server's operating system](#building-from-source).
3.  Copy the compiled `arr-trailer-core` binary and your `config.yaml` file to a directory on your server (e.g., `/mnt/user/appdata/arr-trailer-core/`).
4.  Make the binary executable (one-time command):
    ```bash
    chmod +x arr-trailer-core
    ```
5.  Run the program from the terminal:
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

-   [ ] **Implement Trailer Downloader:** Integrate `yt-dlp` to perform the actual downloads.
-   [ ] **Full Sonarr Integration:** Implement the same detection and downloading logic for TV show episodes.
-   [ ] **Post-Processing:** Add optional steps to call `ffmpeg` for embedding metadata into the downloaded trailer files.
-   [ ] **Caching System:** Implement a local cache to avoid re-scanning and re-querying APIs for media that has been processed recently.
-   [ ] **Language Prioritization:** Allow users to define a preferred language list for trailer searches on TMDB.
-   [ ] **Interactive Setup:** A potential `setup` command to guide users through creating their first `config.yaml`.
