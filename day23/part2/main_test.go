package main

import (
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestPossibleMoves(t *testing.T) {
	startPuz := parseLiteral(strings.Split(strings.TrimSpace(example1), "\n"))
	s1 := parseLiteral(strings.Split(strings.TrimSpace(step1), "\n"))
	if s1.inMotion[keyT{2, 4}] != 0 {
		t.Fatalf("fatal parse: s1.inMotion[%v]=%v, want 0", keyT{2, 4}, s1.inMotion[keyT{2, 4}])
	}
	if s1.inMotion[keyT{6, 4}] != 0 {
		t.Fatalf("fatal parse: s1.inMotion[%v]=%v, want 0", keyT{6, 4}, s1.inMotion[keyT{6, 4}])
	}
	s2 := parseLiteral(strings.Split(strings.TrimSpace(step2), "\n"))
	s3 := parseLiteral(strings.Split(strings.TrimSpace(step3), "\n"))
	s4 := parseLiteral(strings.Split(strings.TrimSpace(step4), "\n"))
	s5 := parseLiteral(strings.Split(strings.TrimSpace(step5), "\n"))
	s6 := parseLiteral(strings.Split(strings.TrimSpace(step6), "\n"))
	s7 := parseLiteral(strings.Split(strings.TrimSpace(step7), "\n"))

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
			name: "blocked D",
			from: keyT{4, 2},
		},
		{
			name: "second blocked A",
			from: keyT{8, 2},
		},
		{
			name: "A moves to left",
			puz:  s1,
			from: keyT{8, 2},
			want: []moveT{
				{from: keyT{8, 2}, to: keyT{7, 0}, energy: 3},
				{from: keyT{8, 2}, to: keyT{9, 0}, energy: 3},
				{from: keyT{8, 2}, to: keyT{5, 0}, energy: 5},
				{from: keyT{8, 2}, to: keyT{3, 0}, energy: 7},
				{from: keyT{8, 2}, to: keyT{1, 0}, energy: 9},
				{from: keyT{8, 2}, to: keyT{0, 0}, energy: 10},
			},
		},
		{
			name: "B moves to right",
			puz:  s2,
			from: keyT{6, 1},
			want: []moveT{
				{from: keyT{6, 1}, to: keyT{5, 0}, energy: 20},
				{from: keyT{6, 1}, to: keyT{7, 0}, energy: 20},
				{from: keyT{6, 1}, to: keyT{3, 0}, energy: 40},
				{from: keyT{6, 1}, to: keyT{9, 0}, energy: 40},
				{from: keyT{6, 1}, to: keyT{1, 0}, energy: 60},
			},
		},
		{
			name: "next B moves to right",
			puz:  s3,
			from: keyT{6, 2},
			want: []moveT{
				{from: keyT{6, 2}, to: keyT{5, 0}, energy: 30},
				{from: keyT{6, 2}, to: keyT{7, 0}, energy: 30},
				{from: keyT{6, 2}, to: keyT{3, 0}, energy: 50},
				{from: keyT{6, 2}, to: keyT{1, 0}, energy: 70},
			},
		},
		{
			name: "next A moves to left",
			puz:  s4,
			from: keyT{6, 3},
			want: []moveT{
				{from: keyT{6, 3}, to: keyT{5, 0}, energy: 4},
				{from: keyT{6, 3}, to: keyT{3, 0}, energy: 6},
				{from: keyT{6, 3}, to: keyT{1, 0}, energy: 8},
			},
		},
		{
			name: "C moves into place",
			puz:  s5,
			from: keyT{4, 1},
			want: []moveT{
				{from: keyT{4, 1}, to: keyT{6, 3}, energy: 600},
			},
		},
		{
			name: "next C moves into place",
			puz:  s6,
			from: keyT{4, 2},
			want: []moveT{
				{from: keyT{4, 2}, to: keyT{6, 2}, energy: 600},
			},
		},
		{
			name: "next B moves into hallway to unblock D",
			puz:  s7,
			from: keyT{4, 3},
			want: []moveT{
				{from: keyT{4, 3}, to: keyT{3, 0}, energy: 40},
				{from: keyT{4, 3}, to: keyT{5, 0}, energy: 40},
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
	want := "Solution: 44169\n"
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
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########
`

var step1 = `
#############
#..........D#
###B#C#B#.###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########
`

var step2 = `
#############
#A.........D#
###B#C#B#.###
  #D#C#B#.#
  #D#B#A#C#
  #A#D#C#A#
  #########
`

var step3 = `
#############
#A........BD#
###B#C#.#.###
  #D#C#B#.#
  #D#B#A#C#
  #A#D#C#A#
  #########
`

var step4 = `
#############
#A......B.BD#
###B#C#.#.###
  #D#C#.#.#
  #D#B#A#C#
  #A#D#C#A#
  #########
`

var step5 = `
#############
#AA.....B.BD#
###B#C#.#.###
  #D#C#.#.#
  #D#B#.#C#
  #A#D#C#A#
  #########
`

var step6 = `
#############
#AA.....B.BD#
###B#.#.#.###
  #D#C#.#.#
  #D#B#C#C#
  #A#D#C#A#
  #########
`

var step7 = `
#############
#AA.....B.BD#
###B#.#.#.###
  #D#.#C#.#
  #D#B#C#C#
  #A#D#C#A#
  #########
`

var step8 = `
#############
#AA...B.B.BD#
###B#.#.#.###
  #D#.#C#.#
  #D#.#C#C#
  #A#D#C#A#
  #########
`

var step9 = `
#############
#AA.D.B.B.BD#
###B#.#.#.###
  #D#.#C#.#
  #D#.#C#C#
  #A#.#C#A#
  #########
`

var step10 = `
#############
#AA.D...B.BD#
###B#.#.#.###
  #D#.#C#.#
  #D#.#C#C#
  #A#B#C#A#
  #########
`

var step11 = `
#############
#AA.D.....BD#
###B#.#.#.###
  #D#.#C#.#
  #D#B#C#C#
  #A#B#C#A#
  #########
`

var step12 = `
#############
#AA.D......D#
###B#.#.#.###
  #D#B#C#.#
  #D#B#C#C#
  #A#B#C#A#
  #########
`

var step13 = `
#############
#AA.D......D#
###B#.#C#.###
  #D#B#C#.#
  #D#B#C#.#
  #A#B#C#A#
  #########
`

var step14 = `
#############
#AA.D.....AD#
###B#.#C#.###
  #D#B#C#.#
  #D#B#C#.#
  #A#B#C#.#
  #########
`

var step15 = `
#############
#AA.......AD#
###B#.#C#.###
  #D#B#C#.#
  #D#B#C#.#
  #A#B#C#D#
  #########
`

var step16 = `
#############
#AA.......AD#
###.#B#C#.###
  #D#B#C#.#
  #D#B#C#.#
  #A#B#C#D#
  #########
`

var step17 = `
#############
#AA.......AD#
###.#B#C#.###
  #.#B#C#.#
  #D#B#C#D#
  #A#B#C#D#
  #########
`

var step18 = `
#############
#AA.D.....AD#
###.#B#C#.###
  #.#B#C#.#
  #.#B#C#D#
  #A#B#C#D#
  #########
`

var step19 = `
#############
#A..D.....AD#
###.#B#C#.###
  #.#B#C#.#
  #A#B#C#D#
  #A#B#C#D#
  #########
`

var step20 = `
#############
#...D.....AD#
###.#B#C#.###
  #A#B#C#.#
  #A#B#C#D#
  #A#B#C#D#
  #########
`

var step21 = `
#############
#.........AD#
###.#B#C#.###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########
`

var step22 = `
#############
#..........D#
###A#B#C#.###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########
`

var step23 = `
#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########
`
