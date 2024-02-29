package scrabble

import "strings"

var SCORES = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) int {
	var score int
	for _, v := range word {
		for key := range SCORES {
			if strings.Contains(key, strings.ToUpper(string(v))) {
				score += SCORES[key]
			}
		}
	}
	return score
}

/*
n: length of the input word
m: the number of keys in the SCORES map
k: length of key in SCORES map assuming string.Contains is O(n)
Time Complexity - O(n*m*k)
Space Complexity - O(1)
*/
