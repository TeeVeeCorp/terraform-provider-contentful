package contentful

import (
	"regexp"
	"strings"
	"unicode"
)

func toSnakeCase(str string) string {
	re := regexp.MustCompile(`[^\w]+`)
	str = re.ReplaceAllString(str, "_")

	var result []rune
	for i, r := range str {
		if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLower(rune(str[i-1])) || (unicode.IsDigit(rune(str[i-1])) && !unicode.IsDigit(r)) {
				result = append(result, '_')
			}
		}
		result = append(result, unicode.ToLower(r))
	}

	finalStr := string(result)
	finalStr = strings.TrimSpace(strings.ReplaceAll(finalStr, "__", "_"))
	finalStr = strings.Trim(finalStr, "_")
	return finalStr
}
