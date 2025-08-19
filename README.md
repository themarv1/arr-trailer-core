# Arr Trailer Core (ATC)

A sleek and blazing-fast command-line tool that synchronizes trailers with your **Radarr** and **Sonarr** instances via their API.

---

### About Arr Trailer Core

**Arr Trailer Core (ATC)** was developed in **Go** to maximize performance and minimize resource consumption. In contrast to complex solutions with a user interface, ATC focuses on what's essential: retrieving movie and series information via API, and then seamlessly and efficiently synchronizing trailers. The downloaded trailers work perfectly in media libraries from **Plex**, **Emby**, and **Jellyfin**.

### Key Features

* **Speed & Efficiency**: As a compiled Go binary, ATC starts extremely quickly and runs with minimal overhead.
* **Headless & CLI-based**: Ideal for automation in Docker or as a scheduled task.
* **YAML Configuration**: Your settings are neatly managed in a single, easy-to-read `.yaml` file.
* **API-Driven Integration**: ATC retrieves information directly from the **Radarr** and **Sonarr** APIs, ensuring your entire library is filled with the right trailers.
