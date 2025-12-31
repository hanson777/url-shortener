package handler

import (
	"net/url"
	"regexp"
	"strings"
)

func normalizeUrl(input string) string {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = "https://" + input
	}
	return input
}

var tldRegex = regexp.MustCompile(`\.[a-zA-Z]{2,}$`)

func isValidURL(input string) bool {
	parsed, err := url.Parse(input)
	if err != nil {
		return false
	}
	return tldRegex.MatchString(parsed.Hostname())
}
