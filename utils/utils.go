package utils

import "strings"

func Reverse(s string) string {

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		temp := r[i]
		r[i] = r[j]
		r[j] = temp
	}

	return string(r)
}

//Splitter splits a string based on the seperators
func Splitter(str string, sep string) []string {
	s := strings.Split(str, sep)

	return s
}
