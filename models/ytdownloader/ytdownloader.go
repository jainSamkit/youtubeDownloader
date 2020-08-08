package ytdownloader

import (
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
	videolinks []types.VideoLink

	//response struct
	Respipe utils.ResponsePipe

	//signatureDecoder
	signaturedecoder signaturedecoder.SignatureDecoder
}

//GetVideoID mines for the video ID from the youtube URL
func (d *Ytdownloader) GetVideoID(url string) string {

	r, _ := regexp.Compile(`v=[\w-]{11}`)

	return r.FindString(url)
}

func (d *Ytdownloader) parseStreamingData() {

	streamingData := d.videolinkinfo.StreamingData
	if streamingData == nil {
		d.Respipe.Success = false
		d.Respipe.SetError("Streaming data is empty.")
		d.Respipe.Info = "No video links found!"

		return
	}

	//collect formats and adaptive formats from the streaming data.
	//The json has url or encrypted signatures to the video files.Also present are the video details such as title and short description.
	formats := streamingData.(map[string]interface{})["formats"]
	adaptiveFormats := streamingData.(map[string]interface{})["adaptiveFormats"]

	videoDetails := d.videolinkinfo.VideoDetails

	allformats := append(formats.([]interface{}), adaptiveFormats.([]interface{})...)

	if len(allformats) == 0 {
		d.Respipe.Success = false
		d.Respipe.SetError("No video formats found!")
		return
	}

	//iterate over the formats and append info to the result
	for k := range allformats {
		var link types.VideoLink
		link.SetInfo(allformats[k].(map[string]interface{}), &(d.signaturedecoder), videoDetails)
		d.videolinks = append(d.videolinks, link)
	}

	d.Respipe.Success = true
	d.Respipe.SetError("None!")
	d.Respipe.Info = "All video links are set!"
}

//New function creates an usable downloader
func New(url string) *Ytdownloader {
	d := Ytdownloader{videoURL: url}
	d.videoID = d.GetVideoID(url)

	d.videopageHTML = ""
	d.signatureJSURL = ""
	return &d
}

//GetVideoLinks gets the links to download video and starts the download.
//If the download fails returns failure
func (d *Ytdownloader) GetVideoLinks() []types.VideoLink {

	if len(d.videoURL) == 0 {
		return d.videolinks
	}
	d.videopageHTML = d.browser.Get(d.videoURL)

	//set the video link information
	d.Respipe = d.videolinkinfo.SetVideoLinkInfo(d.videopageHTML)

	//set the signature js url
	d.signatureJSURL = d.videolinkinfo.GetSignatureJSURL(d.videopageHTML)

	//fetch the js file that contains the signatute encoding info
	// if d.signatureJSURL != "" {
	// 	d.signaturedecoder.SignaturefileJS = d.browser.Get(d.signatureJSURL)
	// 	d.signaturedecoder.ExtractDecoder()
	// 	// fmt.Println("The length of sig is ", len(d.signaturedecoder.SignaturefileJS))
	// 	// directoryname := "C:/Users/samkit jain/Desktop/goprojects/videohtmlfiles/"
	// 	// filename := directoryname + "videojs"
	// 	// utils.WriteinFile(filename, d.signaturedecoder.SignaturefileJS)
	// }

	//extract all the links from the player_response
	d.parseStreamingData()

	//return all the video links
	return d.videolinks
}
