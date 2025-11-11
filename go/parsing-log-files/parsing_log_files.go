package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<\W+>|<>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)"(.*password)"`)

	for _, line := range lines {
		if re.MatchString(line) {
			count += 1
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	response := []string{}
	re := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)

	for _, line := range lines {
		if !re.MatchString(line) {
			response = append(response, line)
			continue
		}

		match := re.FindString(line)
		userString := strings.Fields(match)
		username := userString[len(userString)-1]
		response = append(response, "[USR] "+username+" "+line)
	}

	return response
}
