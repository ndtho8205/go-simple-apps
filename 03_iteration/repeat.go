package iteration

import "strings"

// Repeat returns a string whose value is the concatenation of the given string repeated count times
func Repeat(s string, count int) string {
	// TODO: handle when count < 0
	var repeated string

	for i := 0; i < count; i++ {
		repeated += s
	}

	return repeated
}

func StandardRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}
