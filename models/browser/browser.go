package browser

import (
	"io/ioutil"
	"net/http"
)

//Browser object to send and get requests
type Browser struct {
	storageDir string
	filename   string
}

func (b *Browser) New() *Browser {
	browser := Browser{}
	return &browser
}

//Get function
func (b *Browser) Get(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))

	return string(body)
}

func (b *Browser) GetBytes(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))

	return body
}
