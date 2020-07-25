package downloader

import (
	"fmt"
	"os"

	"github.com/jainSamkit/youtubeDownloader/models/browser"
)

//Downloader stores url information
type Downloader struct {
	VideoURL string
	browser  browser.Browser
}

//New creates the new instance for downloader
func New(url string) *Downloader {
	k := Downloader{VideoURL: url}
	return &k
}

//Download downloads the video in the given directory.
func (d *Downloader) Download() error {

	body := d.browser.GetBytes(d.VideoURL)

	fmt.Println("I am here")
	// if err != nil {
	// 	return err
	// }

	// rep := "movies"
	// os.MkdirAll(rep, 0777)

	filename := "movie1.mp4"
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	fmt.Println(len(body))

	_, err = file.Write(body)
	if err != nil {
		return err
	}

	return nil

}
