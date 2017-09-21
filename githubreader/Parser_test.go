package githubreader

import "testing"

func TestParseGitHubTitle(t *testing.T) {
	cases := []struct {
		title, expected string
	}{
		{"fann", "fann"},
		{"/f/an/n/", "-f-an-n-"},
		{"fa/nn", "fa-nn"},
		{"t3/", "t3-"},
	}

	// verify that everything is correct
	for _, c := range cases {
		title := c.title

		if c.expected != ParseGitHubTitle(title) {
			t.Errorf("Parser(%s) -> '%s', wants '%s'", title, ParseGitHubTitle(title), c.expected)
		}
	}

} // TestParseGitHubTitle
