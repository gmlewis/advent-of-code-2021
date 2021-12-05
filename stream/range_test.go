package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIterate_Plus1(t *testing.T) {
	ch := Iterate(0, func(i, last int) (int, bool) { return last + 1, true })
	got := Take(ch, 10)
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !cmp.Equal(got, want) {
		t.Errorf("Iterate = %+v, want %+v", got, want)
	}
}

func TestIterate_Fib(t *testing.T) {
	ch := Iterate([]int{1, 1}, func(i int, last []int) ([]int, bool) {
		return []int{last[1], last[0] + last[1]}, true
	})
	got := Take(ch, 10)
	want := [][]int{{1, 1}, {1, 2}, {2, 3}, {3, 5}, {5, 8}, {8, 13}, {13, 21}, {21, 34}, {34, 55}, {55, 89}}
	if !cmp.Equal(got, want) {
		t.Errorf("Iterate = %+v, want %+v", got, want)
	}
}

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
			got := ToSlice(Range(tt.start, tt.end))
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Range = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestRanges(t *testing.T) {
	want := [][]int{{0, 3, 0}, {1, 2, 0}, {2, 1, 0}}
	got := ToSlice(Ranges([]int{0, 3, 0}, []int{2, 1, 0}))
	if !cmp.Equal(got, want) {
		t.Errorf("Ranges = %+v, want %+v", got, want)
	}
}
