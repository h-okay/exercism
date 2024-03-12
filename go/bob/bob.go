// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

var REPLACE_LIST = []string{
	" ", "\t", "\n", "\r",
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	s := Cleanup(remark)
	if s == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(s, "?")
	isAllUpper := isAllUpper(s)

	if isQuestion && isAllUpper {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if isAllUpper {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func Cleanup(s string) string {
	for _, v := range REPLACE_LIST {
		s = strings.ReplaceAll(s, string(v), "")
	}
	return s
}

func isAllUpper(s string) bool {
	hasLetters := false
	allUpper := true
	for _, v := range s {
		if unicode.IsLetter(v) {
			hasLetters = true
			if !unicode.IsUpper(v) {
				allUpper = false
			}
		}
	}
	return hasLetters && allUpper
}
