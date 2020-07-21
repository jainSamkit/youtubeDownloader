package utils

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//WriteinFile string body in the filepath provided
func WriteinFile(filepath string, data string) {

	f, err := os.Create(filepath)
	check(err)

	defer f.Close()
	n3, err := f.WriteString(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
}

func WriteBytesinFile(filepath string, data []uint8) {

	f, err := os.Create(filepath)
	check(err)

	defer f.Close()
	n3, err := f.Write(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
}
