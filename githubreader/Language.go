package githubreader

import (
	"encoding/json"
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

	url := langStruct.BaseURL + langStruct.baseURLSuffix

	// extract the languages from the json bytes
	var data map[string]interface{}
	_ = json.Unmarshal(getJSONData(url), &data)

	// add them to a string array
	l := []string{}
	for lang := range data {
		l = append(l, lang)
	}

	langStruct.Language = l
}
