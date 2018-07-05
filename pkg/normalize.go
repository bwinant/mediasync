package pkg

import (
	"strings"
	"regexp"
)

var stopWords = NewSet()
func init() {
	stopWords.Add("a")
	stopWords.Add("an")
	stopWords.Add("and")
	stopWords.Add("the")
}

func Normalize(s string) string {
	// 0) Lower case
	l := strings.ToLower(s)

	// 1) Split on whitespace and path separators
	re1, _ := regexp.Compile(`[\s/\\]`)
	tokens := re1.Split(l, -1)

	// 2) Replace any non alphanumeric chars in each token
	// 3) Remove all stop words
	re2, _ := regexp.Compile(`[^A-Za-z0-9.]`)
	for i, token := range tokens {
		token = re2.ReplaceAllString(token, "")
		if stopWords.Contains(token) {
			token = ""
		}
		tokens[i] = token
	}

	// 4) Join
	return strings.Join(tokens, "")
}