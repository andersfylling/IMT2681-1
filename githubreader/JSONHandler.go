package githubreader

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// getJSONData return a byte array of a json string
func getJSONData(url string) []byte {

	// otherwise we extract the languages
	client := http.Client{
		Timeout: time.Second * 3, // Maximum of 3 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	return body
}
