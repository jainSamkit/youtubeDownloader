package main

import (
	"github.com/jainSamkit/youtubeDownloader/models/ytdownloader"
)

func main() {

	s := "https://www.youtube.com/watch?v=gOMhN-hfMtY"

	d := ytdownloader.New(s)

	d.Downloadvideo()

	// videoID := videoinfo.GetVideoID(s)

	// if len(videoID) > 0 {
	// 	//start downloading videos here
	// 	//call youtube downloader class to download the video

	// 	var a [10]int
	// 	for i := 0; i < len(a); i++ {
	// 		fmt.Println(a[i])
	// 	}
	// } else {
	// 	fmt.Println("No URL found")
	// }

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
