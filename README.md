<p align="center">
  <img src="https://github.com/Superredstone/youtubeDownloaderGo/blob/master/Assets/Download.png" width=200></img> <br><br>
  <img src="https://img.shields.io/github/go-mod/go-version/Superredstone/youtubeDownloaderGo?label=Version&style=for-the-badge">
   <img src="https://img.shields.io/github/languages/top/Superredstone/youtubeDownloaderGo?style=for-the-badge">
  <img src="https://img.shields.io/github/license/Superredstone/youtubeDownloaderGo?color=Green&style=for-the-badge">
  <img src="https://img.shields.io/discord/821836676607115304?label=Discord&style=for-the-badge"></img>
<p>

# Youtube Downloader Go 
Youtube downloader Go is a simple library and application written in Golang for downloading Youtube Videos

# Features 
- Download Youtube Videos
- Easy to use
- Lightweight library

# Application usage
Download clicking [here](https://github.com/Superredstone/youtubeDownloaderGo/releases)

```bash
./YoutubeDownloaderGo-<YOUR OS>-<YOURARCH> -u "https://www.youtube.com/watch?v=aqz-KE-bpKQ" -o "BigBuckBunny"
```

# Getting started
Download library: 
```bash
go get github.com/Superredstone/youtubeDownloaderGo/pkg
```

# Example
```go
package main

import (
    ytDownloader "github.com/Superredstone/youtubeDownloaderGo/pkg"
)

func main() {
    ytDownloader.Download("https://www.youtube.com/watch?v=aqz-KEbpKQ", "BBB.mp4")
}
```

# TODO
  - [ ] Download a video list
  - [ ] Download an entire playlist

# GNU GPL V3
