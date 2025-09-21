package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
        switch char {
            case '❗':
            	return "recommendation"
        	case '🔍':
            	return "search"
        	case '☀':
            	return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    modifiedLog := []rune(log)
	for index, char := range modifiedLog {
        if char == oldRune {
            modifiedLog[index] = newRune
        }
    }
    return string(modifiedLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	logRune := []rune(log)
    if len(logRune) <= limit {
        return true
    } else {
        return false
    }
}
