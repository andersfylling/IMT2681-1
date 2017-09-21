package githubreader

import (
	"encoding/json"
	"log"
)

// Repo The assignment response struct.
// This will hold the payload of our response to a given
// Github repository url
type Repo struct {
	baseURL  string
	Project  string
	Owner    string
	Commiter string
	Commits  int
	Language []string
}

// NewRepo Create a new payload instance
func NewGitHubRepo(username, repository string) *Repo {
	u := ParseGitHubTitle(username)
	r := ParseGitHubTitle(repository)
	return &Repo{
		"https://api.github.com/repos/",
		"github.com/" + u + "/" + r,
		u,
		"",
		0,
		[]string{},
	}
}

// GetLanguages returnes the languages
// It's important to note that you can retrieve the cached results
// or extract fresh results
//
// @param cache true for cache look up
func (repoStruct *Repo) GetRepoDetails(cache bool) {

	commiter := NewCommitor(repoStruct.baseURL)
	commiter.GetCommitor(cache)
	repoStruct.Commiter = commiter.Username
	repoStruct.Commits = commiter.Commits

	languages := NewLanguages(repoStruct.baseURL)
	languages.GetLanguages(cache)
	repoStruct.Language = languages.Language

}

// GetJSON Returns a JSON string of the object
func (ghr *Repo) GetJSON() string {
	obj, err := json.Marshal(ghr)

	if err != nil {
		log.Fatal(err)
		return "{}"
	}

	return string(obj)
}
