package logs

import "unicode/utf8"

const defaultApp = "default"

var appList = map[rune]string{
    '❗': "recommendation",
    '🔍': "search",
    '☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
        if name, ok := appList[c]; ok {
            return name
        }
    }
    return defaultApp
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    runes := []rune(log)
    for i, r := range runes {
    	if r == oldRune {
    		runes[i] = newRune
    	}
    }
    return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
