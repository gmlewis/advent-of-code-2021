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

	want := "Solution: 4512\n"
	if got != want {
		t.Errorf("process = %q, want %q", got, want)
	}
}

var got string

func testPrintf(format string, a ...interface{}) (int, error) {
	got = fmt.Sprintf(format, a...)
	return 0, nil
}

var example1 = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`
