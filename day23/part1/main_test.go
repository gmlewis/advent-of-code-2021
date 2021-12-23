package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestClearPath(t *testing.T) {
}

func TestEnergy(t *testing.T) {
}

func TestPossibleMoves(t *testing.T) {
	startPuz := parse(strings.Split(strings.TrimSpace(example1), "\n"))

	tests := []struct {
		name         string
		puz          *puzT
		from         keyT
		want         []keyT
		wantEnergies []int
	}{
		{
			name:         "first B",
			from:         keyT{2, 1},
			want:         []keyT{{1, 0}, {3, 0}, {0, 0}, {5, 0}, {7, 0}, {9, 0}, {10, 0}},
			wantEnergies: []int{20, 20, 30, 40, 60, 80, 90},
		},
		{
			name:         "first C",
			from:         keyT{4, 1},
			want:         []keyT{{3, 0}, {5, 0}, {1, 0}, {7, 0}, {0, 0}, {9, 0}, {10, 0}},
			wantEnergies: []int{200, 200, 400, 400, 500, 600, 700},
		},
		{
			name:         "second B",
			from:         keyT{6, 1},
			want:         []keyT{{5, 0}, {7, 0}, {3, 0}, {9, 0}, {10, 0}, {1, 0}, {0, 0}},
			wantEnergies: []int{20, 20, 40, 40, 50, 60, 70},
		},
		{
			name:         "first D",
			from:         keyT{8, 1},
			want:         []keyT{{7, 0}, {9, 0}, {10, 0}, {5, 0}, {3, 0}, {1, 0}, {0, 0}},
			wantEnergies: []int{2000, 2000, 3000, 4000, 6000, 8000, 9000},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.puz == nil {
				tt.puz = startPuz
			}
			got, energies := tt.puz.possibleMoves(tt.from)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("possibleMoves(%v) = keys %#v, want %#v", tt.from, got, tt.want)
			}
			if !cmp.Equal(energies, tt.wantEnergies) {
				t.Errorf("possibleMoves(%v) = energies %#v, want %#v", tt.from, energies, tt.wantEnergies)
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
