package gitHub

import (
	"bytes"
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

// charCode Get the ASCII code for the given rune
func charCode(c rune) int {
	return int(c)
}

// ParseGithubTitle Parses a title by the Github rules.
// I don't actually know the rules, I just noticed some runes/chars
// that they accept.
func ParseGitHubTitle(title string) string {
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
