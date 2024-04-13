package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var output strings.Builder

	for _, v := range plain {
		if unicode.IsPunct(v) || unicode.IsDigit(v) {
			output.WriteString(string(v))
			continue
		}

		if unicode.IsSpace(v) {
			output.WriteString(" ")
			continue
		}

		if unicode.IsUpper(v) {
			output.WriteString(strings.ToUpper(shift(v, shiftKey)))
			continue
		}

		if unicode.IsLower(v) {
			output.WriteString(shift(v, shiftKey))
			continue
		}
	}

	return output.String()
}

func shift(char rune, shiftKey int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	start := strings.Index(alphabet, strings.ToLower(string(char)))
	return string(alphabet[(start+shiftKey)%len(alphabet)])
}
