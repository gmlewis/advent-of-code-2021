package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 9\n"
	test.Runner(t, example1, want, process, &printf)
}

var example1 = `
3,4,3,1,2
`
