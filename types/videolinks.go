package types

import (
	"net/url"
	"regexp"

	"github.com/jainSamkit/youtubeDownloader/models/signaturedecoder"
	"github.com/jainSamkit/youtubeDownloader/utils"
)

//VideoLink structure stores refined url and tags for all the videos
type VideoLink struct {
	VideoURL    string
	VideoTag    float64
	VideoFormat string
	VideoTitle  string
}

//SetInfo function to set url from the video map and also decrypts signature
func (link *VideoLink) SetInfo(videoMap map[string]interface{}, sigdecoder *signaturedecoder.SignatureDecoder, videoDetails interface{}) {

	videourl := videoMap["url"]
	itag := videoMap["itag"]

	link.VideoTag = itag.(float64)
	link.VideoFormat = utils.GetTagInfo(link.VideoTag)
	link.VideoTitle = videoDetails.(map[string]interface{})["title"].(string)

	var signature string
	if videourl != nil {
		link.VideoURL = videourl.(string)
		r, _ := regexp.Compile(`u0026`)
		e := r.ReplaceAllString(link.VideoURL, `&`)
		link.VideoURL = e

	} else {

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

			signature = parsedCipher["s"][0]
			sigdecoder.Signature = signature
			signature = sigdecoder.DecodeSignature()

			url := parsedCipher["url"][0]
			sp := parsedCipher["sp"][0]
			link.VideoURL = url + "&" + sp + "=" + signature
		}
	}

}
