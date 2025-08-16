package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
    pattern := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return pattern.MatchString(text)
}

func SplitLogLine(text string) []string {
	separator := regexp.MustCompile(`<[~*=\-]*>`)
    return separator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(s string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(s, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	result := make([]string, len(lines))

	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			username := match[1]
			result[i] = "[USR] " + username + " " + line
		} else {
			result[i] = line
		}
	}
	return result
}
