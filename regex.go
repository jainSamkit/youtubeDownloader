package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := `"player_res":`

	r, _ := regexp.Compile(`.*`)

	fmt.Println(r.FindString(s))

}
