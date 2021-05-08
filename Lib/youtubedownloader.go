package ytDownloader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Download(url, output string) error {
	videoURL := getVideoPosition(url)

	if videoURL == "Error" {
		fmt.Println("Error, unable to fetch APIs")
	}

	err := downloadFile(videoURL, output)

	fmt.Println("Downloaded file")
	if err != nil {
		fmt.Println("Unable to download file")
		return err
	}

	return nil
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
	if file.Error == "No links found" {
		return "Error"
	} else {
		return file.Link[0]
	}
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return err
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

type API struct {
	Link  []string `json:"links"`
	Error string   `json:"error"`
}
