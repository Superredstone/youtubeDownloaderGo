package main

import (
	"flag"

	ytDownloader "github.com/Superredstone/youtubeDownloaderGo/Lib"
)

func main() {
	url := flag.String("u", "", "Youtube video URL")
	out := flag.String("o", "", "Output name")

	flag.Parse()

	ytDownloader.Download(*url, *out+".mp4")
}
