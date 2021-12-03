package must

// FilterFunc takes a string and returns true if the
// string is to be kept.
type FilterFunc func(string) bool

// FilterStrings filters a slice of strings and keeps
// those values for which f returns true.
func FilterStrings(lines []string, f FilterFunc) []string {
	var result []string
	for _, line := range lines {
		if f(line) {
			result = append(result, line)
		}
	}
	return result
}
