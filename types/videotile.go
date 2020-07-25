package types

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/jainSamkit/youtubeDownloader/models/browser"
	"github.com/jainSamkit/youtubeDownloader/models/signaturedecoder"
	"github.com/jainSamkit/youtubeDownloader/utils"
)

//VideoTile structure stores refined url and tags for all the videos
type VideoTile struct {
	videoURL    string
	videoTag    float64
	videoFormat string
	browser     browser.Browser
}

//SetInfo function to set url from the video map and also decrypts signature
func (tile *VideoTile) SetInfo(videoMap map[string]interface{}, sigdecoder *signaturedecoder.SignatureDecoder) {

	videourl := videoMap["url"]
	itag := videoMap["itag"]

	tile.videoTag = itag.(float64)
	tile.videoFormat = utils.GetTagInfo(tile.videoTag)

	if videourl != nil {
		tile.videoURL = videourl.(string)
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

		signature := parsedCipher["s"][0]
		// signature := "8=g87CIbR1kjWPWi7mKaPoV0wUXN-hT0t571GGE=lSyljAiAwxYaXG79Knn4aLSum6jBqhMe1gD35jd8QhlAjDxVx7NAhIQRw8JQ0qOq"

		sigdecoder.Signature = signature

		fmt.Println(signature)
		signature = sigdecoder.DecodeSignature()
		fmt.Println(signature)

		fmt.Println("")
		url := parsedCipher["url"][0]
		sp := parsedCipher["sp"][0]

		// 'url' => $url . '&' . $sp . '=' . $decoded_signature,

		tile.videoURL = url + "&" + sp + "=" + signature

	}
}
