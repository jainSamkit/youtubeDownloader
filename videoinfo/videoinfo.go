package videoinfo

import (
	"regexp"
	"strings"
)

//GetVideID mines for the video ID from the youtube URL
func GetVideoID(url string) string {

	r, _ := regexp.Compile(`v=[\w]{11}`)

	return r.FindString(url)
}

//GetPlayerResponse mines out the player_response json that contains links to videos
func GetPlayerResponse(s string) string {
	r, _ := regexp.Compile(`player_response":"(.*?)","`)
	s = r.FindString(s)
	if len(s) > 0 {
		s = strings.ReplaceAll(s, `\`, ``)
	}

	return s
}
