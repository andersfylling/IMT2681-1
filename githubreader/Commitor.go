package githubreader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Commitor The assignment response struct.
// This will hold the payload of our response to a given
// Github repository url
type Commitor struct {
	position      int // 0 = contributor with most commits
	baseURLSuffix string
	BaseURL       string
	Username      string
	Commits       int
}

// NewCommitor Create a new payload instance
func NewCommitor(url string) *Commitor {
	return &Commitor{
		0,
		"/contributors",
		url,
		"",
		0,
	}
}

// GetCommitor returnes the languages
// It's important to note that you can retrieve the cached results
// or extract fresh results
//
// @param cache true for cache look up
func (comStruct *Commitor) GetCommitor(cache bool) {

	if cache && len(comStruct.Username) > 0 {
		return
	}

	// otherwise we extract the languages
	client := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
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
	var data []struct {
		Login   string `json:"login"`
		Commits int    `json:"contributions"`
	}
	_ = json.Unmarshal(body, &data)

	comStruct.Username = data[comStruct.position].Login
	comStruct.Commits = data[comStruct.position].Commits

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
