package main

import (
	"fmt"

	"github.com/jainSamkit/youtubeDownloader/videoinfo"
)

func main() {

	s := "https://www.youtube.com/watch?v=zEmtfA8FETc"

	videoID := videoinfo.GetVideoID(s)

	fmt.Println(videoID)

	// // fmt.Println(s)
	// // k := utils.Splitter(s, "?")
	// // fmt.Println(s)
	// // fmt.Println(k)

	// // matched, err := regexp.MatchString(`^a.*d`, "avbfdd")
	// // fmt.Println(matched)
	// // fmt.Println(err)

	// // a := "axabbb"
	// r, _ := regexp.Compile(`v=[\w]{11}`)
	// fmt.Println(r.FindString(s))

	// fmt.Printf("%T\n", r.FindString(s))

	// // r2, _ := regexp.Compile(`^a.*d$`)
	// r3, _ := regexp.Compile(`(\w+ \d+)`)
	// a := `Jan 1987`
	// // b := `w14x`

	// fmt.Println(r3.FindAllString(a, 2))
}
