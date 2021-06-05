package main

import (
	"flag"
	"fmt"
	"os"

	ytDownloader "github.com/Superredstone/youtubeDownloaderGo/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("./YoutubeDownloaderGo -u <Video URL> -o <FileName>")
		os.Exit(1)
	}

	url := flag.String("u", "", "Youtube video URL")
	out := flag.String("o", "", "Output name")

	if *out == "" {
		*out = "Download"
	}

	flag.Parse()

	err := ytDownloader.Download(*url, *out+".mp4")
	if err != nil {
		fmt.Println(err)
	}
}
