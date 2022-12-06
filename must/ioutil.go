package must

import (
	"os"
	"strings"
)

// ReadFile reads a file as a string or causes a fatal error.
// It trims extra trailing whitespace from the string.
func ReadFile(filename string) string {
	buf, err := os.ReadFile(filename)
	if err != nil {
		fatal(err)
	}

	return strings.TrimRight(string(buf), " \t\n")
}

// ReadFileLines reads a file and returns a slice of lines.
func ReadFileLines(filename string) []string {
	return ReadSplitFile(filename, "\n")
}

// ReadSplitFile reads a file and splits on the given string.
func ReadSplitFile(filename, split string) []string {
	s := ReadFile(filename)
	return strings.Split(s, split)
}
