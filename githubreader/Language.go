package githubreader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Languages The assignment response struct.
// This will hold the payload of our response to a given
// Github repository url
type Languages struct {
	baseURLSuffix string
	BaseURL       string
	Language      []string
}

// NewLanguages Create a new payload instance
// @param url should be the base url for the api repository
func NewLanguages(url string) *Languages {
	return &Languages{
		"/languages",
		url,
		[]string{},
	}
}

// GetLanguages returnes the languages
// It's important to note that you can retrieve the cached results
// or extract fresh results
//
// @param cache true for cache look up
func (langStruct *Languages) GetLanguages(cache bool) {

	// If languages has already been set
	if cache && len(langStruct.Language) > 0 {
		return
	}

	// otherwise we extract the languages
	client := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, langStruct.BaseURL+langStruct.baseURLSuffix, nil)
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

	// add them to a string array
	l := []string{}
	for lang := range data {
		l = append(l, lang)
	}

	langStruct.Language = l
}

// GetJSON Returns a JSON string of the object
// func (ghr *Repo) GetJSON() string {
// 	obj, err := json.Marshal(ghr)
//
// 	if err != nil {
// 		log.Fatal(err)
// 		return "{}"
// 	}
//
// 	return string(obj)
// }
