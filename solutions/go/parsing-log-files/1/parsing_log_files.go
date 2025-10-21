package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	var validPrefixes = []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}

	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"[^"]*[pP][aA][sS][sS][wW][oO][rR][dD][^"]*"`)
	var count int
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		count += len(matches)
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	var taggedLines []string
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			username := matches[1]
			taggedLines = append(taggedLines, "[USR] "+username+" "+line)
		} else {
			taggedLines = append(taggedLines, line)
		}
	}
	return taggedLines
}
