package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.WriteString(example1); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	printf = testPrintf
	process(tmpfile.Name())

	want := "Solution: 7\n"
	if got != want {
		t.Errorf("process = %q, want %q", got, want)
	}
}

var got string

func testPrintf(format string, a ...interface{}) (int, error) {
	got = fmt.Sprintf(format, a...)
	return 0, nil
}

var example1 = `199
200
208
210
200
207
240
269
260
263`
