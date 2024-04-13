package isbn

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {
	sanitized := sanitize(isbn)
	if len(sanitized) != 10 || invalidCheckChar(sanitized) {
		return false
	}

	total := 0
	multiplier := 1
	for i := len(sanitized) - 1; i >= 0; i-- {
		digit := sanitized[i]

		if digit == 'X' {
			total += 10 * multiplier
			multiplier++
			continue
		}

		total += int(digit-'0') * multiplier
		multiplier++
	}
	
	return total%11 == 0
}

func sanitize(s string) string {
	pattern := `\w+`
	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllString(s, -1)
	return strings.Join(matches, "")
}

func invalidCheckChar(s string) bool {
	return strings.ContainsRune(s, 'X') && len(s)-1 != strings.Index(s, "X")
}
