package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 0\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########


#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########


#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########


#############
#...B.......#
###B#C#.#D###
  #A#D#C#A#
  #########


#############
#...B.......#
###B#.#C#D###
  #A#D#C#A#
  #########


#############
#.....D.....#
###B#.#C#D###
  #A#B#C#A#
  #########


#############
#.....D.....#
###.#B#C#D###
  #A#B#C#A#
  #########


#############
#.....D.D.A.#
###.#B#C#.###
  #A#B#C#.#
  #########


#############
#.........A.#
###.#B#C#D###
  #A#B#C#D#
  #########


#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########

`
