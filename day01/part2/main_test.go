package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 5\n"
	test.Runner(t, example1, want, process, &printf)
}

var example1 = `
199
200
208
210
200
207
240
269
260
263
`
