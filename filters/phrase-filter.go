package filters

import (
	"strings"
)

func In(test, alphabet string) bool {
	for w := range alphabet {
		if string(alphabet[w]) == test {
			return true
		}
	}
	return false
}

func PhraseFilter(text string) []FilterResult {

	newText := ""

	for t := range text {
		if !In(string(text[t]), "!£$%&/()=?*[]#@§_-:.;,'") {
			newText = newText + string(text[t])
		}
	}

	words := strings.Split(newText, " ")

	out := MultiWordFilter(words)
	return out

}
