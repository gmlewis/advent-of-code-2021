package strfn

// Equals returns a function that checks if a string
// is equal to value.
func Equals(value string) func(string) bool {
	return func(s string) bool { return s == value }
}

// Length returns the length of the provided string.
func Length(s string) int { return len(s) }

// Substr returns a function that returns a substring
// from [start, end) (just like s[start:end]).
//
// If start >= len(s), "" is returned.
// If end >= len(s), s[start:] is returned.
func Substr(start, end int) func(string) string {
	return func(s string) string {
		if start >= len(s) {
			return ""
		}
		if end >= len(s) {
			return s[start:]
		}
		return s[start:end]
	}
}

// First returns the first character (as a string) of the provided string
// or "".
func First(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[0:1]
}

// Last returns the last character (as a string) of the provided string
// or "".
func Last(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1:]
}

// RunesWithIndex iterates over a string and calls the provided
// function with its index and rune. This is because I couldn't
// figure out how to make WithIndex work with a string and its runes.
func RunesWithIndex(s string, f func(i int, value rune)) {
	for i, v := range s {
		f(i, v)
	}
}
