package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordCounts := Frequency{}

	wordRegex := regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z0-9]+)?`)
	words := wordRegex.FindAllString(phrase, -1)

	for _, word := range words {
		word = strings.ToLower(word)
		wordCounts[word]++
	}

	return wordCounts

}
