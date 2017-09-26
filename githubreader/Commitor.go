package githubreader

import (
	"encoding/json"
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

	url := comStruct.BaseURL + comStruct.baseURLSuffix

	// extract the languages from the json bytes
	var data []struct {
		Login   string `json:"login"`
		Commits int    `json:"contributions"`
	}
	_ = json.Unmarshal(getJSONData(url), &data)

	comStruct.Username = data[comStruct.position].Login
	comStruct.Commits = data[comStruct.position].Commits
}
