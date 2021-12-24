package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestInput(t *testing.T) {
	tests := []struct {
		name   string
		digits [14]int64
		want   int64
	}{
		{
			name:   "all 9s",
			digits: [14]int64{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			want:   9919490,
		},
		{
			name:   "all 1s",
			digits: [14]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:   6117450,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := input(tt.digits); got != tt.want {
				t.Errorf("input() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
inp x
mul x -1


inp z
inp x
mul z 3
eql z x


inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2

`
