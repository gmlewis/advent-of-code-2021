package main

import (
	"strings"
	"testing"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestEnhance(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(example1), "\n")
	filter := ReduceWithIndex([]rune(lines[0]), filterT{}, func(i int, r rune, acc filterT) filterT {
		if r == '#' {
			acc[i] = 1
		}
		return acc
	})

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "enhance - 1st step",
			in:   strings.TrimSpace(enhance1),
			want: strings.TrimSpace(enhance2),
		},
		{
			name: "enhance - 2nd step",
			in:   strings.TrimSpace(enhance2),
			want: strings.TrimSpace(enhance3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := parse(strings.Split(tt.in, "\n"))
			got := img.enhance(filter, 0).String()
			if got != tt.want {
				t.Errorf("enhance=\n%v\n, want:\n%v", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 35\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###
`

var enhance1 = `
#..#.
#....
##..#
..#..
..###
`

var enhance2 = `
.##.##..
#..#.#..
##.#..#.
####..#.
.#..##..
..##..#.
...#.#..
`

var enhance3 = `
.......#...
.#..#.#....
#.#...###..
#...##.#...
#.....#.#..
.#.#####...
..#.#####..
...##.##...
....###....
`
