// Package test provides some handy testing utilities.
package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

type logFunc *func(string, ...interface{})
type printFunc *func(string, ...interface{}) (int, error)

// Benchmark benchmarks the provided process.
func Benchmark(b *testing.B, filename string, process func(string), logf logFunc, printf printFunc) {
	b.Helper()

	wd, err := os.Getwd()
	if err != nil {
		b.Fatal(err)
	}
	path := filepath.Join(wd, filename)
	*logf = silentLogf
	*printf = silentPrintf

	for n := 0; n < b.N; n++ {
		process(path)
	}
}

// Runner runs the provided process by placing the puzzle
// input into a temporary file and passing the filename to
// the process. It then tests that the resulting output
// (by overriding the `printf` method) matches the desired
// output.
func Runner(t *testing.T, puzzle, want string, process func(string), printf *func(string, ...interface{}) (int, error)) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "example1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(puzzle); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	*printf = testPrintf

	process(tmpfile.Name())

	if got != want {
		t.Errorf("process = %q, want %q", got, want)
	}
}

var got string

func testPrintf(format string, a ...interface{}) (int, error) {
	got = fmt.Sprintf(format, a...)
	return 0, nil
}

func silentLogf(format string, a ...interface{}) {}

func silentPrintf(format string, a ...interface{}) (int, error) { return 0, nil }
