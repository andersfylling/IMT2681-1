package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Setup what runes are legal
func setLegalRunes() [255]bool {
	allowed := [255]bool{}

	// 0 - 9 chars
	for i := 48; i < 58; i++ {
		allowed[i] = true
	}

	// A - Z
	for i := 65; i < 91; i++ {
		allowed[i] = true
	}

	// a - z
	for i := 97; i < 123; i++ {
		allowed[i] = true
	}

	// -
	allowed[45] = true

	// .
	allowed[46] = true

	// _
	allowed[95] = true

	return allowed
}

// http://www.asciitable.com/
// legal runes that's accepted by Github usernames / repository titles
var legal = setLegalRunes()

type repositoryData struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type repositoryDataContributor struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

/*
- Response payload:
{
    "project": {
        "type": "string"
    },
    "owner": {
        "type": "string"
    },
    "committer": {
        "type": "string"
    },

    "commits": {
        "type": "number"
    },
    "language": {
        "array": {
            "items": {
                "type": "string"
            }
        }
    }
}
*/
type assignmentResponse struct {
	Project   string
	Owner     string
	Committer string
	Commits   int
	Language  []string
}

// charCode Get the ASCII code for the given rune
func charCode(c rune) int {
	return int(c)
}

// parseGithubTitle Parses a title by the Github rules.
// I don't actually know the rules, I just noticed some runes/chars
// that they accept.
func parseGithubTitle(title string) string {
	// for every none legal char, we convert to a `.`
	// for every none legal fixes, we don't stack them.
	var buffer bytes.Buffer
	var usedFix bool
	for _, r := range title {
		ascii := charCode(r)

		// if the ascii value isn't within the aSCII table
		// convert its value into a rune we know is viewed as illegal: `!`
		if ascii > 255 || ascii < 0 {
			ascii = 33
		}

		// adds the rune, but format it if it's illegal
		if !legal[ascii] {
			// we hit a illegal character, so we need to convert it to `-`
			if !usedFix {
				// If the ruen before this rune was illegal, we don't add another `-`
				// as this isn't done on the Github site.
				buffer.WriteString("-")
				usedFix = true
			}
		} else {
			// a legal ASCII char/rune was hit, it is then added to the buffer.
			buffer.WriteRune(r)
			usedFix = false
		}
	}

	return buffer.String()
}

// GenerateGitHubRepositoryURL Takes a username and repository to generate a psuedo legal api url
func GenerateGitHubRepositoryURL(username string, repository string) string {
	var buffer bytes.Buffer

	// concate the api url for the repository
	buffer.WriteString("https://api.github.com/repos/")
	buffer.WriteString(parseGithubTitle(username))
	buffer.WriteString("/")
	buffer.WriteString(parseGithubTitle(repository))

	return buffer.String() // convert to a string
}

func parseRepositoryJSON(w http.ResponseWriter, url string, username string, repository string) {
	data := repositoryData{}
	data2 := []repositoryDataContributor{}
	var data3 map[string]interface{}

	{
		client := http.Client{
			Timeout: time.Second * 2, // Maximum of 2 secs
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

		_ = json.Unmarshal(body, &data)
	}

	{
		client := http.Client{
			Timeout: time.Second * 2, // Maximum of 2 secs
		}

		req, err := http.NewRequest(http.MethodGet, url+"/contributors", nil)
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

		_ = json.Unmarshal(body, &data2)
	}

	{
		client := http.Client{
			Timeout: time.Second * 2, // Maximum of 2 secs
		}

		req, err := http.NewRequest(http.MethodGet, url+"/languages", nil)
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

		_ = json.Unmarshal(body, &data3)
	}

	languages := []string{}
	for lang := range data3 {
		languages = append(languages, lang)
	}

	// create the response
	r := assignmentResponse{
		"github.com/" + username + "/" + repository,
		username,
		data2[0].Login,         // /contributors[0]
		data2[0].Contributions, // /contributors[0].contributions
		languages,              // /languages
	}

	j, err := json.Marshal(r)

	if err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}

	fmt.Fprintf(w, string(j))
}

// ParseGithubAPI Given a username and a repository this returns info about such
// Rules:
//  - The :username must exist
//  - The :repository given, must exist as a property of the :username
func ParseGithubAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url := GenerateGitHubRepositoryURL(ps.ByName("username"), ps.ByName("repository"))
	parseRepositoryJSON(w, url, ps.ByName("username"), ps.ByName("repository"))
}
