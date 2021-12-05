package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Sums: [7 5 8 7 5], gamma=22, toggle=31, epsilon=9, product: 198\n"
	test.Runner(t, example1, want, process, &printf)
}

var example1 = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`
