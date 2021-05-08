package main

import (
	"flag"

	downloader "github.com/Superredstone/youtubeDownloaderGo"
)

func main() {
	var url = flag.String("u", "", "Youtube Video URL")
	var output = flag.String("o", "video.mp4", "Output name")

	flag.Parse()

	downloader.youtubeDownloaderGo.download(url, output)
}
