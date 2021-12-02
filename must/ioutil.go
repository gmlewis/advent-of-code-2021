package must

import (
	"io/ioutil"
	"log"
	"strings"
)

// ReadFile reads a file or causes a fatal error.
// It also returns a trimmed string instead of a byte slice.
func ReadFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(buf))
}

// ReadFileLines reads a file and returns a slice of lines.
func ReadFileLines(filename string) []string {
	s := ReadFile(filename)
	return strings.Split(s, "\n")
}
