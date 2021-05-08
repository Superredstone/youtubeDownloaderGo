package ytDownloader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Download(url, output string) {
	downloadFile(getVideoPosition(url), "Downloads/"+output)
}

func getVideoPosition(ytUrl string) string {
	url := "https://maadhav-ytdl.herokuapp.com/video_info.php?url=" + ytUrl

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var file API
	json.Unmarshal(responseData, &file)

	return file.Link[0]
}

func downloadFile(URL, fileName string) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println(err)
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//Write the bytes to the fiel
	counter := &WriteCounter{}
	_, err = io.Copy(file, io.TeeReader(response.Body, counter))
	if err != nil {
		fmt.Println(err)
	}
}

type API struct {
	Link []string `json:"links"`
}

/////////////////////////////////////////////////////////////Download Bar//////////////////////////////////////////////////////////////

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", wc.Total)
}
