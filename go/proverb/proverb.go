package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	output := make([]string, 0, len(rhyme)+1)
	if len(rhyme) == 0 {
		return output
	}

	for i := 0; i < len(rhyme)-1; i++ {
		sentence := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		output = append(output, sentence)
	}

	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return output
}
