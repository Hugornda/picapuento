package utils

import (
	"fmt"
	"regexp"
)

func ExtractToken(html string) (string, error) {
	re := regexp.MustCompile(`<input type="hidden" value="([^"]*)" id="token" />`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("token not found")
}
