package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 37\n"
	test.Runner(t, example1, want, process, &printf)
}

var example1 = `
16,1,2,0,4,2,7,1,2,14
`
