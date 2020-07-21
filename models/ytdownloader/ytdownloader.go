package ytdownloader

import (
	"encoding/json"
	"fmt"

	"github.com/jainSamkit/youtubeDownloader/models/browser"
	"github.com/jainSamkit/youtubeDownloader/utils"
	"github.com/jainSamkit/youtubeDownloader/videoinfo"
)

//Ytdownloader is a skeleton for the downloader object
type Ytdownloader struct {
	videoID       string
	url           string
	browser       browser.Browser
	videopageHTML string
	playerRes     string
}

//New function creates an usable downloader
func New(url string) *Ytdownloader {
	d := Ytdownloader{url: url}
	d.videoID = videoinfo.GetVideoID(url)

	return &d
}

//Downloadvideo gets the links to download video and starts the download.
//If the download fails returns failure
func (d *Ytdownloader) Downloadvideo() {
	d.videopageHTML = d.browser.Get(d.url)

	//write the html content in a file
	directoryname := "C:/Users/samkit jain/Desktop/goprojects/videohtmlfiles/"
	filename := directoryname + d.videoID
	utils.WriteinFile(filename, d.videopageHTML)

	d.playerRes = videoinfo.GetPlayerResponse(d.videopageHTML)
	strB, _ := json.Marshal(d.playerRes)
	fmt.Printf("The type of strb is %T\n", strB)

	fmt.Println("ok")
	filename += "video_repsonse"
	// utils.WriteinFile(filename, strB)
	utils.WriteBytesinFile(filename, strB)
	// utils.WriteinFile(filename, d.playerRes)
}
