package must

import (
	"os"
	"strings"
)

// ReadFile reads a file as a string or causes a fatal error.
func ReadFile(filename string) string {
	buf, err := os.ReadFile(filename)
	if err != nil {
		fatal(err)
	}

	return string(buf)
}

// ReadFileLines reads a file and returns a slice of lines.
func ReadFileLines(filename string) []string {
	s := ReadFile(filename)
	return strings.Split(s, "\n")
}
