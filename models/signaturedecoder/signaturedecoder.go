package signaturedecoder

import (
	"regexp"
	"strconv"
	"strings"
)

//SignatureDecoder extracts out the decrypter from the js file and decodes the signature passed to it
type SignatureDecoder struct {

	//complete encoder function
	//function(a){a=a.split("");Yu.QC(a,39);Yu.oN(a,77);Yu.QC(a,20);return a.join("")}
	decoderFunction string

	//stores the action of each individual functions
	//{QC:"reverse","oN":"swap", ....}
	funcActionMap map[string]string

	//stores the functions occurring in the same order as in the encoder function
	//[[Yu.QC(a,39); Yu QC a 39] [Yu.oN(a,77); Yu oN a 77] [Yu.QC(a,20); Yu QC a 20]]
	funcWheel [][]string

	//signature js file in string
	SignaturefileJS string
	//signature value
	Signature string
}

func reverse(s string) string {

	// fmt.Println(s)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func swap(s string, value string) string {
	i, _ := strconv.Atoi(value)

	runes := []rune(s)
	temp := runes[0]

	runes[0] = runes[i%(len(s))]
	runes[i%(len(s))] = temp

	return string(runes)
}

func splice(s string, val string) string {

	value, _ := strconv.Atoi(val)
	runes := []rune(s)
	runes = runes[value:]
	return string(runes)
}

//DecodeSignature decodes the signature
func (sigD *SignatureDecoder) DecodeSignature() string {

	decodedString := sigD.Signature

	for i := range sigD.funcWheel {
		funcName := sigD.funcWheel[i][2]
		value := sigD.funcWheel[i][4]

		action := sigD.funcActionMap[funcName]

		if action == "reverse" {
			decodedString = reverse(decodedString)
		} else if action == "swap" {
			decodedString = swap(decodedString, value)
		} else if action == "splice" {
			decodedString = splice(decodedString, value)
		}
	}

	return decodedString
}

//ExtractDecoder extracts out the func required to encode the signature
func (sigD *SignatureDecoder) ExtractDecoder() {

	//function(a){a=a.split("");Yu.QC(a,39);Yu.oN(a,77);Yu.QC(a,20);return a.join("")}
	//extract the above function from the file
	r1, _ := regexp.Compile(`function\(\s*a\s*\)\s*{\s*a\s*=\s*a\.split\(\s*""\s*\)(.*?)return(.*?)}`)
	sigD.decoderFunction = r1.FindString(sigD.SignaturefileJS)

	//extracts individual functions in order
	r2, _ := regexp.Compile(`(\w+)\.(\w+)\((\w+),(\d+)\);`)
	sigD.funcWheel = r2.FindAllStringSubmatch(sigD.decoderFunction, -1)

	sigD.funcActionMap = make(map[string]string)

	// extracts the function definition corresponding to the action
	// ["oN":"reverse","QC":"swap",...]
	for i := range sigD.funcWheel {
		funcName := sigD.funcWheel[i][2]

		// oN:function(a){a.reverse()}
		r3, _ := regexp.Compile(funcName + `\s*:\s*function\s*\(.*?\)\s*{(.*?)}`)
		action := r3.FindString(sigD.SignaturefileJS)

		if strings.Contains(action, "reverse") {
			sigD.funcActionMap[funcName] = "reverse"
		} else if strings.Contains(action, "length") {
			sigD.funcActionMap[funcName] = "swap"
		} else if strings.Contains(action, "splice") {
			sigD.funcActionMap[funcName] = "splice"
		}
	}
}
