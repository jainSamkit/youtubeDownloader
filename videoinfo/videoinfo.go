package videoinfo

import (
	"fmt"
	"regexp"
)

func GetVideoID(url string) string {
	fmt.Println(url)
	r, _ := regexp.Compile(`v=[\w]{11}`)

	return r.FindString(url)
}
