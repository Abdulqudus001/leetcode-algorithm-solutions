package decodestring

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	// regex will only match brackets that have only letters inside
	// i.e innermost brackets
	regex := regexp.MustCompile(`((\d+)\[([a-z]*)\])`)
	matches := regex.FindAllStringSubmatch(s, -1)
	// matches will catch non nested groups and run replacement
	for _, match := range matches {
		if len(match) >= 3 {
			mult, _ := strconv.Atoi(match[2])
			letters := match[3]
			// repeat string, mult times
			replacement := strings.Repeat(letters, mult)
			group := match[0]
			// replace exact group
			s = strings.Replace(s, group, replacement, -1)

		}
	}
	fmt.Println(s)
	// Recurse function to keep solving for outer brackets
	if regex.MatchString(s) {
		s = decodeString(s)
	}
	return s
}
