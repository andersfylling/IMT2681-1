package main

import "testing"

func TestGenerateGitHubRepositoryURL(t *testing.T) {
	cases := []struct {
		username, repository, expected string
	}{
		{"sciencefyll", "fann", "https://api.github.com/repos/sciencefyll/fann"},
		{"sciencefyll", "/f/an/n/", "https://api.github.com/repos/sciencefyll/-f-an-n-"},
		{"sciencefyll", "fa/nn", "https://api.github.com/repos/sciencefyll/fa-nn"},
		{"sciencefyll", "t3/", "https://api.github.com/repos/sciencefyll/t3-"},
	}

	// verify that everything is correct
	for _, c := range cases {
		u := c.username
		r := c.repository

		if c.expected != GenerateGitHubRepositoryURL(u, r) {
			t.Errorf("GenerateGitHubRepositoryURL(%s, %s) -> '%s', expected '%s'", u, r, GenerateGitHubRepositoryURL(u, r), c.expected)
		}
	}

} // TestGenerateGitHubRepositoryURL
