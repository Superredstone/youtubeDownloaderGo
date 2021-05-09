<p align="center">
  <img src="https://github.com/Superredstone/youtubeDownloaderGo/blob/master/Assets/Download.png" width=200></img> <br><br>
  <img src="https://img.shields.io/github/go-mod/go-version/Superredstone/youtubeDownloaderGo?label=Version&style=for-the-badge">
   <img src="https://img.shields.io/github/languages/top/Superredstone/youtubeDownloaderGo?style=for-the-badge">
  <img src="https://img.shields.io/github/license/Superredstone/youtubeDownloaderGo?color=Green&style=for-the-badge">
  <img src="https://img.shields.io/discord/821836676607115304?label=Discord&style=for-the-badge"></img>
<p>

# Youtube Downloader Go
Simple Youtube downloader library and application written in Go

# Getting started with library
```bash
go get github.com/Superredstone/youtubeDownloaderGo/Lib
```

Code example: 
```go
package main

import (
	ytDownloader "github.com/Superredstone/youtubeDownloaderGo/Lib"
)

func main() {
	ytDownloader.Download("https://www.youtube.com/watch?v=aqz-KE-bpKQ", "BigBuckBunny.mp4")
}

```

# Application usage
Test command: 
```./YoutubeDownloaderGo-<your os>-<your aarch> -u "https://www.youtube.com/watch?v=aqz-KE-bpKQ" -o "BigBuckBunny"```
