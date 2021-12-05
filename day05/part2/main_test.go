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

	want := "Solution: 12\n"
	if got != want {
		t.Errorf("process = %q, want %q", got, want)
	}
}

var got string

func testPrintf(format string, a ...interface{}) (int, error) {
	got = fmt.Sprintf(format, a...)
	return 0, nil
}

var example1 = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`
