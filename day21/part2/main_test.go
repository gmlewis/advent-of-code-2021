package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/v1/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 444356092776315\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
Player 1 starting position: 4
Player 2 starting position: 8

`
