package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	sum := 0
	count := 0
	for i := len(id) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(id[i])) {
			return false
		}

		digit := int(id[i] - '0')
		if count%2 != 0 {
			doubled := digit * 2
			if doubled > 9 {
				doubled -= 9
			}
			sum += doubled
			count++
			continue
		}
		sum += digit
		count++
	}

	return sum%10 == 0
}
