package main

import (
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestClearPath(t *testing.T) {
}

func TestEnergy(t *testing.T) {
}

func TestPossibleMoves(t *testing.T) {
	startPuz := parse(strings.Split(strings.TrimSpace(example1), "\n"))
	bMoved := parse(strings.Split(strings.TrimSpace(bMovedExample), "\n"))

	tests := []struct {
		name string
		puz  *puzT
		from keyT
		want []moveT
	}{
		{
			name: "first B",
			from: keyT{2, 1},
			want: []moveT{
				{from: keyT{2, 1}, to: keyT{1, 0}, energy: 20},
				{from: keyT{2, 1}, to: keyT{3, 0}, energy: 20},
				{from: keyT{2, 1}, to: keyT{0, 0}, energy: 30},
				{from: keyT{2, 1}, to: keyT{5, 0}, energy: 40},
				{from: keyT{2, 1}, to: keyT{7, 0}, energy: 60},
				{from: keyT{2, 1}, to: keyT{9, 0}, energy: 80},
				{from: keyT{2, 1}, to: keyT{10, 0}, energy: 90},
			},
		},
		{
			name: "first C",
			from: keyT{4, 1},
			want: []moveT{
				{from: keyT{4, 1}, to: keyT{3, 0}, energy: 200},
				{from: keyT{4, 1}, to: keyT{5, 0}, energy: 200},
				{from: keyT{4, 1}, to: keyT{1, 0}, energy: 400},
				{from: keyT{4, 1}, to: keyT{7, 0}, energy: 400},
				{from: keyT{4, 1}, to: keyT{0, 0}, energy: 500},
				{from: keyT{4, 1}, to: keyT{9, 0}, energy: 600},
				{from: keyT{4, 1}, to: keyT{10, 0}, energy: 700},
			},
		},
		{
			name: "second B",
			from: keyT{6, 1},
			want: []moveT{
				{from: keyT{6, 1}, to: keyT{5, 0}, energy: 20},
				{from: keyT{6, 1}, to: keyT{7, 0}, energy: 20},
				{from: keyT{6, 1}, to: keyT{3, 0}, energy: 40},
				{from: keyT{6, 1}, to: keyT{9, 0}, energy: 40},
				{from: keyT{6, 1}, to: keyT{10, 0}, energy: 50},
				{from: keyT{6, 1}, to: keyT{1, 0}, energy: 60},
				{from: keyT{6, 1}, to: keyT{0, 0}, energy: 70},
			},
		},
		{
			name: "first D",
			from: keyT{8, 1},
			want: []moveT{
				{from: keyT{8, 1}, to: keyT{7, 0}, energy: 2000},
				{from: keyT{8, 1}, to: keyT{9, 0}, energy: 2000},
				{from: keyT{8, 1}, to: keyT{10, 0}, energy: 3000},
				{from: keyT{8, 1}, to: keyT{5, 0}, energy: 4000},
				{from: keyT{8, 1}, to: keyT{3, 0}, energy: 6000},
				{from: keyT{8, 1}, to: keyT{1, 0}, energy: 8000},
				{from: keyT{8, 1}, to: keyT{0, 0}, energy: 9000},
			},
		},
		{
			name: "first blocked A",
			from: keyT{2, 2},
		},
		{
			name: "blocked D",
			from: keyT{4, 2},
		},
		{
			name: "blocked C",
			from: keyT{6, 2},
		},
		{
			name: "second blocked A",
			from: keyT{8, 2},
		},
		{
			name: "C moves into place",
			puz:  bMoved,
			from: keyT{4, 1},
			want: []moveT{
				{from: keyT{4, 1}, to: keyT{6, 1}, energy: 400},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.puz == nil {
				tt.puz = startPuz
			}
			got := tt.puz.possibleMoves(tt.from)
			if len(got) != len(tt.want) {
				t.Errorf("possibleMoves(%v) = %#v, want %#v", tt.from, got, tt.want)
			}
			for i := range got {
				if got[i].from != tt.want[i].from || got[i].to != tt.want[i].to || got[i].energy != tt.want[i].energy {
					t.Errorf("possibleMoves(%v)[%v] = %#v, want %#v", tt.from, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 12521\n"
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
`

var bMovedExample = `
#############
#...B.......#
###B#C#.#D###
  #A#D#C#A#
  #########
`
