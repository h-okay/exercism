package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	subjectCounter := getCounter(subject)
	output := make([]string, 0, len(candidates))

	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}

		counter := getCounter(candidate)
		if reflect.DeepEqual(subjectCounter, counter) {
			output = append(output, candidate)
		}
	}

	return output
}

func getCounter(str string) map[rune]int {
	counter := make(map[rune]int)
	for _, ch := range str {
		counter[unicode.ToLower(ch)]++
	}
	return counter
}
