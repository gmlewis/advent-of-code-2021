package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/v1/test"
)

func TestShoot(t *testing.T) {
	tests := []struct {
		name      string
		steps     int
		wantXVMin int
		wantYVMin int
		wantXVMax int
		wantYVMax int
	}{
		{
			name:      "1 step",
			steps:     1,
			wantXVMin: 20,
			wantYVMin: -10,
			wantXVMax: 30,
			wantYVMax: -5,
		},
		{
			name:      "2 steps",
			steps:     2,
			wantXVMin: 11,
			wantYVMin: -5,
			wantXVMax: 15,
			wantYVMax: -2,
		},
		{
			name:      "first example: 7 steps",
			steps:     7,
			wantXVMin: 6,
			wantYVMin: 2,
			wantXVMax: 7,
			wantYVMax: 2,
		},
		{
			name:      "second example: 9 steps",
			steps:     9,
			wantXVMin: 6,
			wantYVMin: 3,
			wantXVMax: 7,
			wantYVMax: 3,
		},
		{
			name:      "third example: 4 steps",
			steps:     4,
			wantXVMin: 7,
			wantYVMin: -1,
			wantXVMax: 9,
			wantYVMax: 0,
		},
		{
			name:      "winning example: 20 steps",
			steps:     20,
			wantXVMin: 6,
			wantYVMin: 9,
			wantXVMax: 7,
			wantYVMax: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXVMin, gotYVMin, gotXVMax, gotYVMax := shoot(tt.steps, 20, -10, 30, -5)
			if gotXVMin != tt.wantXVMin {
				t.Errorf("shoot() gotXVMin = %v, want %v", gotXVMin, tt.wantXVMin)
			}
			if gotYVMin != tt.wantYVMin {
				t.Errorf("shoot() gotYVMin = %v, want %v", gotYVMin, tt.wantYVMin)
			}
			if gotXVMax != tt.wantXVMax {
				t.Errorf("shoot() gotXVMax = %v, want %v", gotXVMax, tt.wantXVMax)
			}
			if gotYVMax != tt.wantYVMax {
				t.Errorf("shoot() gotYVMax = %v, want %v", gotYVMax, tt.wantYVMax)
			}
		})
	}
}

func TestCalcHeight(t *testing.T) {
	tests := []struct {
		name   string
		steps  int
		want   int
		wantOK bool
	}{
		{
			name:   "first example: 7,2",
			steps:  7,
			want:   3,
			wantOK: true,
		},
		{
			name:   "second example: 6,3",
			steps:  9,
			want:   6,
			wantOK: true,
		},
		{
			name:   "third example: 9,0",
			steps:  4,
			want:   0,
			wantOK: true,
		},
		{
			name:   "winning example: 6,9",
			steps:  20,
			want:   45,
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := calcHeight(tt.steps, 20, 30, -10, -5, map[keyT]bool{})
			if gotOK != tt.wantOK {
				t.Fatalf("simulate = %v, want %v", gotOK, tt.wantOK)
			}
			if got != tt.want {
				t.Fatalf("simulate = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulate(t *testing.T) {
	tests := []struct {
		name   string
		xv     int
		yv     int
		want   int
		wantOK bool
	}{
		{
			name:   "first example: 7,2",
			xv:     7,
			yv:     2,
			want:   3,
			wantOK: true,
		},
		{
			name:   "second example: 6,3",
			xv:     6,
			yv:     3,
			want:   6,
			wantOK: true,
		},
		{
			name:   "third example: 9,0",
			xv:     9,
			yv:     0,
			want:   0,
			wantOK: true,
		},
		{
			name:   "fourth example: 17,-4",
			xv:     17,
			yv:     -4,
			want:   0,
			wantOK: false,
		},
		{
			name:   "winning example: 6,9",
			xv:     6,
			yv:     9,
			want:   45,
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := simulate(tt.xv, tt.yv, 20, 30, -10, -5)
			if gotOK != tt.wantOK {
				t.Fatalf("simulate = %v, want %v", gotOK, tt.wantOK)
			}
			if got != tt.want {
				t.Fatalf("simulate = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 45\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
target area: x=20..30, y=-10..-5
`
