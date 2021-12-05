package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name  string
		start int
		end   int
		want  []int
	}{
		{
			name: "no steps",
			want: []int{0},
		},
		{
			name:  "one step, ascending",
			start: 10,
			end:   11,
			want:  []int{10, 11},
		},
		{
			name:  "one step, descending",
			start: 110,
			end:   109,
			want:  []int{110, 109},
		},
		{
			name:  "ten steps, ascending",
			start: 0,
			end:   9,
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "ten steps, ascending",
			start: 9,
			end:   0,
			want:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Range(tt.start, tt.end)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Range = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestRanges(t *testing.T) {
	tests := []struct {
		name  string
		start []int
		end   []int
		want  [][]int
	}{
		{
			name: "no steps",
			want: [][]int{{}},
		},
		{
			name:  "45 degrees",
			start: []int{0, 3, 0},
			end:   []int{2, 1, 0},
			want:  [][]int{{0, 3, 0}, {1, 2, 0}, {2, 1, 0}},
		},
		{
			name:  "fractional increment",
			start: []int{0, 3, 0},
			end:   []int{5, 1, 0},
			want:  [][]int{{0, 3, 0}, {1, 3, 0}, {2, 2, 0}, {3, 2, 0}, {4, 1, 0}, {5, 1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ranges(tt.start, tt.end)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Ranges = %+v, want %+v", got, tt.want)
			}
		})
	}
}
