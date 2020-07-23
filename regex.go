package main

import (
	"fmt"
)

func main() {
	// s := `"player_res":"{\"reponse_ctx\":{},\"play_status\":{},\"streaming_data\":{"formats":[{}]}}",
	// "player_ads":"{\"reponse_ctx\":{},\"play_status\":{},\"streaming_data\":{"formats":[{}]}}",
	// "player_ad1":"{\"reponse_ctx\":{},\"play_status\":{},\"streaming_data\":{"formats":[{}]}}"`

	// r, _ := regexp.Compile(`"player_res":`)

	// fmt.Println(r.FindString())

	var i interface{}

	i = 2011
	i = "232"
	i = 2.777

	fmt.Printf("The type of i is%T\n", i)

}
