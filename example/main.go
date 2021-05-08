package main

import (
	"flag"
)

func main() {
	var url = flag.String("u", "", "Youtube Video URL")
	var output = flag.String("o", "video.mp4", "Output name")

	flag.Parse()

	youtubeDownloaderGo.youtubeDownloaderGo.download(url, output)
}
