package main

import (
	"fmt"

	"github.com/jainSamkit/youtubeDownloader/models/ytdownloader"
)

func main() {

	// //url already set.
	// // s := "https://www.youtube.com/watch?v=n1JUAR8q3LU"

	// // s := "https://www.youtube.com/watch?v=Uw5JOtvFd-k&t=1s"

	// //json error
	// s := "https://www.youtube.com/watch?v=k9zTr2MAFRg"

	// s := "https://www.youtube.com/watch?v=ldZdJJxS0m8"

	s := "https://www.youtube.com/watch?v=1I-3vJSC-Vo"

	// //url not set,signature set
	// s := "https://www.youtube.com/watch?v=lFGnsdV-sR4"

	// // s := "https://www.youtube.com/watch?v=2xDnxkzQtdI"
	d := ytdownloader.New(s)

	videolinks := d.GetVideoLinks()

	fmt.Println(len(videolinks))

	// var videoURL string

	// // fmt.Println(videolinks)
	// for i := range videolinks {
	// 	if strings.Contains(videolinks[i].VideoFormat, "video") && strings.Contains(videolinks[i].VideoFormat, "360p") &&
	// 		strings.Contains(videolinks[i].VideoFormat, "mp4") {
	// 		videoURL = videolinks[i].VideoURL
	// 		break
	// 	}
	// }

	// // videoURL = "https://r1---sn-gwpa-25ur.googlevideo.com/videoplayback?expire=1595721605u&ei=JXMcX7_DJfT-4-EPn-Ci2AIu&ip=2409%3A4064%3Ab81%3A1be0%3A7933%3Abbb4%3A4667%3Ad84u&id=o-AOQhdVKlCmcNiQqoR737zNy86DEVDB7TvqtzQboYJtDLu&itag=18u&source=youtubeu&requiressl=yesu&mh=hLu&mm=31%2C29u&mn=sn-gwpa-25ur%2Csn-qxa7sn7lu&ms=au%2Crduu&mv=mu&mvi=1u&pl=36u&initcwndbps=100000u&vprv=1u&mime=video%2Fmp4u&gir=yesu&clen=8148879u&ratebypass=yesu&dur=104.025u&lmt=1484577245498715u&mt=1595699846u&fvip=1u&fexp=23883098u&beids=23886208u&c=WEBu&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cgir%2Cclen%2Cratebypass%2Cdur%2Clmtu&sig=AOq0QJ8wRgIhAPb8GwgwvzE_wFYarymjSn43AlHkYusPQmdXzUH6kyY_AiEArHsprqTu0Us2G-rS_rKU2BVoiqfK47CVEXn3sfu8tCI%3Du&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbpsu&lsig=AG3C_xAwRAIgevLQu7cU3-Pq5gy4DOVZDCZBwucsex6PmTHJzD19LM8CIBhqcIDfStx05O_XIQ0kbj8p46bneQyFOTxMUK5kYIoh"

	// fmt.Println(videoURL)
	// dwd := downloader.New(videoURL)
	// err := dwd.Download()
	// fmt.Println(err)
}
