package main

import (
	"flag"
	"fmt"

	ytDownloader "github.com/Superredstone/youtubeDownloaderGo/Lib"
)

func main() {
	url := flag.String("u", "", "Youtube video URL")
	out := flag.String("o", "", "Output name")

	flag.Parse()

	err := ytDownloader.Download(*url, *out+".mp4")
	if err != nil {
		fmt.Println(err)
	}
}
