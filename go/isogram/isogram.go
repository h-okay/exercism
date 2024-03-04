package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	w := strings.ToLower(word)
	for idx, char := range strings.ToLower(word) {
		if unicode.IsLetter(char) && strings.ContainsRune(w[idx+1:], char) {
			return false
		}
	}
	return true

}
