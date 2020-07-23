package types

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/jainSamkit/youtubeDownloader/utils"
)

//VideoTile structure stores refined url and tags for all the videos
type VideoTile struct {
	videoURL    string
	videoTag    float64
	videoFormat string
}

//SetInfo function to set url from the video map and also decrypts signature
func (tile *VideoTile) SetInfo(videoMap map[string]interface{}, signatureJS string) {

	videourl := videoMap["url"]
	itag := videoMap["itag"]

	if videourl != nil {
		tile.videoURL = videourl.(string)
		tile.videoTag = itag.(float64)
		tile.videoFormat = utils.GetTagInfo(tile.videoTag)
		return
	}

	//write the function parser for decrypting the signature
	signatureCipher := videoMap["signatureCipher"]
	cipher := videoMap["cipher"]

	cipherKey := signatureCipher
	if cipherKey == nil {
		cipherKey = cipher
	}

	if cipherKey != nil {

		r, _ := regexp.Compile(`u0026`)
		e := r.ReplaceAllString(cipherKey.(string), `&`)

		parsedCipher, err := url.ParseQuery(e)
		if err != nil {
			return
		}

		fmt.Println(parsedCipher)
	}
}
