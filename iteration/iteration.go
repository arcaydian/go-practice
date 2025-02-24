package iteration

import "strings"

// Repated takes a character which is repeated repeatTimes
func Repeated(character string, repeatTimes int8) string {
	var repeated strings.Builder
	for range repeatTimes {
		repeated.WriteString(character)
	}
	return repeated.String()
}
