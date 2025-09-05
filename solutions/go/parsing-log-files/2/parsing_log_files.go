package parsinglogfiles

import "regexp"

var validLinePattern = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var separatorPattern = regexp.MustCompile(`<[~*=\-]*>`)
var passwordPattern = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
var endOfLinePattern = regexp.MustCompile(`end-of-line\d+`)
var userPattern = regexp.MustCompile(`User\s+(\S+)`)

func IsValidLine(text string) bool {
	return validLinePattern.MatchString(text)
}

func SplitLogLine(text string) []string {
    return separatorPattern.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if passwordPattern.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(s string) string {
	return endOfLinePattern.ReplaceAllString(s, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		match := userPattern.FindStringSubmatch(line)
		if len(match) > 1 {
			username := match[1]
			result[i] = "[USR] " + username + " " + line
		} else {
			result[i] = line
		}
	}
	return result
}
