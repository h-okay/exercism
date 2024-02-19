package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	output := strings.Repeat("*", numStarsPerLine)
	output += "\n"
	output += welcomeMsg
	output += "\n"
	output += strings.Repeat("*", numStarsPerLine)
	return output
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	output := strings.Replace(oldMsg, "*", "", -1)
	output = strings.Trim(output, " \n")
	return output
}
