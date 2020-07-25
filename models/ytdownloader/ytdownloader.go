package ytdownloader

import (
	"fmt"
	"regexp"

	"github.com/jainSamkit/youtubeDownloader/models/browser"
	"github.com/jainSamkit/youtubeDownloader/models/signaturedecoder"
	"github.com/jainSamkit/youtubeDownloader/types"
	"github.com/jainSamkit/youtubeDownloader/utils"
)

//Ytdownloader is a skeleton for the downloader object
type Ytdownloader struct {
	videoID        string
	videoURL       string
	browser        browser.Browser
	videopageHTML  string
	signatureJSURL string

	//videolinkinfo to process the streaming data
	videolinkinfo types.VideoLinkInfo

	//collection of videos with url,tag and video format(mp4...)
	videotiles []types.VideoTile

	//response struct
	resPipe utils.ResponsePipe

	//signatureDecoder
	signaturedecoder signaturedecoder.SignatureDecoder
}

//GetVideoID mines for the video ID from the youtube URL
func (d *Ytdownloader) GetVideoID(url string) string {

	r, _ := regexp.Compile(`v=[\w]{11}`)

	return r.FindString(url)
}

func (d *Ytdownloader) parseStreamingData() {

	streamingData := d.videolinkinfo.StreamingData

	formats := streamingData.(map[string]interface{})["formats"]
	adaptiveFormats := streamingData.(map[string]interface{})["adaptiveFormats"]

	allformats := append(formats.([]interface{}), adaptiveFormats.([]interface{})...)

	if len(allformats) == 0 {
		d.resPipe.Success = false
		d.resPipe.SetError("No video formats found!")
		return
	}

	fmt.Println("The length of formats is ", len(allformats))

	//iterate over the formats and append info to the result

	for k := range allformats {
		var tile types.VideoTile
		tile.SetInfo(allformats[k].(map[string]interface{}), &(d.signaturedecoder))
		d.videotiles = append(d.videotiles, tile)
	}

	d.resPipe.Success = true
	d.resPipe.SetError("None!")
	d.resPipe.Info = "All video tiles are set!"
}

//New function creates an usable downloader
func New(url string) *Ytdownloader {
	d := Ytdownloader{videoURL: url}
	d.videoID = d.GetVideoID(url)
	d.videopageHTML = ""
	d.signatureJSURL = ""
	return &d
}

//Downloadvideo gets the links to download video and starts the download.
//If the download fails returns failure
func (d *Ytdownloader) Downloadvideo() {
	if len(d.videoURL) == 0 {
		// print("No url found!")
		return
	}
	d.videopageHTML = d.browser.Get(d.videoURL)

	//set the video link information
	d.videolinkinfo.SetVideoLinkInfo(d.videopageHTML)

	//set the signature js url
	d.signatureJSURL = d.videolinkinfo.GetSignatureJSURL(d.videopageHTML)

	//fetch the js file that contains the signatute encoding info
	if d.signatureJSURL != "" {
		d.signaturedecoder.SignaturefileJS = d.browser.Get(d.signatureJSURL)
		d.signaturedecoder.ExtractDecoder()
		// fmt.Println("The length of sig is ", len(d.signaturedecoder.SignaturefileJS))
		// directoryname := "C:/Users/samkit jain/Desktop/goprojects/videohtmlfiles/"
		// filename := directoryname + "videojs"
		// utils.WriteinFile(filename, d.signaturedecoder.SignaturefileJS)
	}

	//extract all the links from the
	d.parseStreamingData()
}
