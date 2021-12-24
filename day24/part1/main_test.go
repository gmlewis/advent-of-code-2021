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
			name:   "all 1s",
			digits: [14]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:   6117450,
		},
		{
			name:   "all 2s",
			digits: [14]int64{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			want:   6592705,
		},
		{
			name:   "all 3s",
			digits: [14]int64{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
			want:   7067960,
		},
		{
			name:   "all 4s",
			digits: [14]int64{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
			want:   7543215,
		},
		{
			name:   "all 5s",
			digits: [14]int64{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
			want:   8018470,
		},
		{
			name:   "all 6s",
			digits: [14]int64{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
			want:   8493725,
		},
		{
			name:   "all 7s",
			digits: [14]int64{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
			want:   8968980,
		},
		{
			name:   "all 8s",
			digits: [14]int64{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
			want:   9444235,
		},
		{
			name:   "all 9s",
			digits: [14]int64{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			want:   9919490,
		},
		{
			name:   "all 9s with digits[6]-1==digits[7]",
			digits: [14]int64{9, 9, 9, 9, 9, 9, 9, 8, 9, 9, 9, 9, 9, 9},
			want:   381518,
		},
		{
			name:   "all 9s with digits[6]-1==digits[7] and digits[11]-2==digits[12]",
			digits: [14]int64{9, 9, 9, 9, 9, 9, 9, 8, 9, 9, 9, 9, 7, 9},
			want:   381512,
		},
		{
			name:   "all 9s with digits[6]-1==digits[7] and digits[11]-2==digits[12] and digits[10]+5==digits[11]",
			digits: [14]int64{9, 9, 9, 9, 9, 9, 9, 8, 9, 9, 4, 9, 7, 9},
			want:   381512,
		},
		{
			name: `all 9s with:
               digits[6]-1==digits[7] and
               digits[11]-2==digits[12] and
               digits[10]+5==digits[11] and
               digits[9]+2==digits[10]`,
			digits: [14]int64{9, 9, 9, 9, 9, 9, 9, 8, 9, 2, 4, 9, 7, 9},
			want:   14678,
		},
		{
			name: `all 9s with:
		           digits[6]-1==digits[7] and
		           digits[11]-2==digits[12] and
		           digits[10]+5==digits[11] and
		           digits[9]+2==digits[10] and
		           digits[2]!=digits[3]`,
			digits: [14]int64{9, 9, 9, 8, 9, 9, 9, 8, 9, 2, 4, 9, 7, 9},
			want:   381564,
		},
		{
			name: `all 9s with:
               digits[6]-1==digits[7] and
               digits[11]-2==digits[12] and
               digits[10]+5!=digits[11] and
               digits[9]+2==digits[10] and
               digits[8]+2==digits[9]`,
			digits: [14]int64{9, 9, 9, 9, 9, 9, 9, 8, 1, 3, 5, 9, 7, 9},
			want:   14678,
		},
		{
			name: `all 9s with:
               digits[6]-1==digits[7] and
               digits[11]-2==digits[12] and
               digits[10]+5!=digits[11] and
               digits[9]+2==digits[10] and
               digits[8]+2==digits[9] and
               digits[5]+3==digits[10]`,
			digits: [14]int64{9, 9, 9, 9, 9, 2, 9, 8, 1, 3, 5, 9, 7, 9},
			want:   560,
		},
		{
			name: `all 9s with:
               digits[6]-1==digits[7] and
               digits[11]-2==digits[12] and
               digits[10]+5!=digits[11] and
               digits[9]+2==digits[10] and
               digits[8]+2==digits[9] and
               digits[5]+3==digits[10] and
               digits[4]-5==digits[11] and`,
			digits: [14]int64{9, 9, 9, 9, 9, 2, 9, 8, 1, 3, 5, 4, 2, 9},
			want:   14,
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
