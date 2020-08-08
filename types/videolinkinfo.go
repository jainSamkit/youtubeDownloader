package types

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/jainSamkit/youtubeDownloader/utils"
)

//VideoLinkInfo structure to hold all the video link related information
type VideoLinkInfo struct {
	VideolinkJSON string
	StreamingData interface{}
	VideoDetails  interface{}
}

//GetVideoLinkJSON selects the player response from the page html.
func (v *VideoLinkInfo) GetVideoLinkJSON(videohtml string) string {

	// fmt.Println(videohtml)
	r, _ := regexp.Compile(`"player_response":"(.*?)","`)
	json := r.FindString(videohtml)

	json = strings.ReplaceAll(json, `\`, ``)
	json = strings.Replace(json, `"player_response":"`, `"player_response":`, 1)
	r2, _ := regexp.Compile(`codecs="(.*?)"`)
	json = r2.ReplaceAllString(json, ``)

	// r3, _ := regexp.Compile(`"(.*)"[\w-_.]"(.*)"`)
	// json = r3.FindString(json)

	// fmt.Println("Printing json")
	// fmt.Println(json)

	jsonrune := []rune(json)

	//to remove "," from the end of the json string
	jsonrune = jsonrune[:len(jsonrune)-3]

	json = string(jsonrune)

	json += string("}")
	json = string("{") + json

	return json
}

//RemoveOffending removes any offending string of type(" simpleText: "Hello "Go" !")
func (v *VideoLinkInfo) RemoveOffending(bytejson []byte, offendChar string) ([]byte, bool) {

	// "simpleText": "Just a guy from the end of the world making Motivational videos to inspire everyone who is looking for some "
	// motivation " or advice on life."

	//Below function finds : "Just a guy from the end of the world making Motivational videos to inspire everyone who is looking for some "motivation"
	r, _ := regexp.Compile(`:\s*"` + `[^"]*?` + `("` + offendChar + `[^"]*?")`)

	offendingString := r.FindStringSubmatch(string(bytejson))

	if len(offendingString) > 1 {
		completeString := string(offendingString[0])
		fatalString := string(offendingString[1])
		replaceString := strings.ReplaceAll(fatalString, `"`, ``)
		finalString := strings.ReplaceAll(completeString, fatalString, replaceString)
		modifiedString := r.ReplaceAllString(string(bytejson), finalString)

		// directoryname := "C:/Users/samkit jain/Desktop/goprojects/videohtmlfiles/"
		// filename := directoryname + "modified_json"
		// utils.WriteinFile(filename, modifiedString)

		return []byte(modifiedString), false
	} else {
		return bytejson, true
	}
}

//SetVideoLinkInfo mines out the player_response json and stores it in a map
func (v *VideoLinkInfo) SetVideoLinkInfo(videohtml string) utils.ResponsePipe {

	var resPipe utils.ResponsePipe

	v.VideolinkJSON = v.GetVideoLinkJSON(videohtml)

	bytejson := []byte(v.VideolinkJSON)

	// directoryname := "C:/Users/samkit jain/Desktop/goprojects/videohtmlfiles/"
	// filename := directoryname + "player_res_json"
	// utils.WriteBytesinFile(filename, bytejson)

	var raw map[string]interface{}

	//unmarshall bytejson to raw structure

	err := json.Unmarshal(bytejson, &raw)

	for err != nil {
		r4, _ := regexp.Compile(`invalid character '(.*)'`)
		errorJSON := r4.FindStringSubmatch(err.Error())
		offendChar := errorJSON[1]
		bytejson, isBreak := v.RemoveOffending(bytejson, offendChar)
		err = json.Unmarshal(bytejson, &raw)
		if isBreak {
			break
		}
	}

	if err == nil {
		//Heirarchy of streamingData ==> player_response-->responseContext-->streamingData.
		//Streaming data holds all the information about links inside formats and adaptive formats field.
		linkData := raw["player_response"].(interface{})
		linkData = linkData.(map[string]interface{})["streamingData"]

		videodetail := raw["player_response"].(interface{})
		videodetail = videodetail.(map[string]interface{})["videoDetails"]

		v.VideoDetails = videodetail
		v.StreamingData = linkData
		resPipe.Success = true
	} else {
		resPipe.Success = false
	}

	return resPipe
}

//GetSignatureJSURL sets the loc of sig function js file
func (v *VideoLinkInfo) GetSignatureJSURL(videohtml string) string {

	//extract the script url
	r, _ := regexp.Compile(`<script\s*src="[^"]+player[^"]+js"`)
	scriptURL := r.FindString(videohtml)

	//trim <script src part
	r1, _ := regexp.Compile(`<script\s*src=`)
	scriptURL = r1.ReplaceAllString(scriptURL, ``)

	//trim all the tabs and whitespaces
	r2, _ := regexp.Compile(`\s+`)
	scriptURL = r2.ReplaceAllString(scriptURL, "")

	//trim all the "
	r3, _ := regexp.Compile(`"`)
	scriptURL = r3.ReplaceAllString(scriptURL, "")

	var sigjsURL string
	//check for relative url scheme
	if strings.Index(scriptURL, "//") == 0 {
		sigjsURL = "http:" + scriptURL
	} else if strings.Index(scriptURL, "/") == 0 {
		sigjsURL = "http://www.youtube.com" + scriptURL
	} else {
		sigjsURL = ""
	}

	return sigjsURL
}
