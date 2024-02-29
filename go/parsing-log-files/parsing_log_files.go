package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	valid_prefixes := []string{"TRC", "DBG", "INF", "WRN", "ERR", "FTL"}
	pattern := fmt.Sprintf(`^\[(%s)\]`, strings.Join(valid_prefixes, "|"))
	re, _ := regexp.Compile(pattern)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, len(lines))
	re, _ := regexp.Compile(`User\s+(\S+)`)

	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match == nil {
			taggedLines[i] = line
			continue
		}
		userName := match[1]
		taggedLines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
	}

	return taggedLines

}
