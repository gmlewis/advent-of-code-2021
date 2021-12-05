package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: [15 10] - product: 150\n"
	test.Runner(t, example1, want, process, &printf)
}

var example1 = `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`
