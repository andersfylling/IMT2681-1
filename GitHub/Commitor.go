package GitHub

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Repo The assignment response struct.
// This will hold the payload of our response to a given
// Github repository url
type Commitor struct {
	position      int // 0 = contributor with most commits
	baseURLSuffix string
	BaseURL       string
	Username      string
	Commits       int
}

// NewRepo Create a new payload instance
func NewCommitor(url string) *Commitor {
	return &Commitor{
		0,
		"/contributors",
		url,
		"",
		0,
	}
}

// GetLanguages returnes the languages
// It's important to note that you can retrieve the cached results
// or extract fresh results
//
// @param cache true for cache look up
func (comStruct *Commitor) GetCommitor(cache bool) {

	// otherwise we extract the languages
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, comStruct.BaseURL+comStruct.baseURLSuffix, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// extract the languages from the json bytes
	var data map[string]interface{}
	_ = json.Unmarshal(body, &data)

}

// GetJSON Returns a JSON string of the object
// func (c *Commitor) GetJSON() string {
// 	obj, err := json.Marshal(c)
//
// 	if err != nil {
// 		log.Fatal(err)
// 		return "{}"
// 	}
//
// 	return string(obj)
// }
