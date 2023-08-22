package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIterate_Plus1(t *testing.T) {
	t.Parallel()
	ch := Iterate(0, func(i, last int) (int, bool) { return last + 1, true })
	got := Take(ch, 10)
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !cmp.Equal(got, want) {
		t.Errorf("Iterate = %+v, want %+v", got, want)
	}
}

func TestIterate_Fib(t *testing.T) {
	t.Parallel()
	ch := Iterate([]int{1, 1}, func(i int, last []int) ([]int, bool) {
		return []int{last[1], last[0] + last[1]}, true
	})
	got := Take(ch, 10)
	want := [][]int{{1, 1}, {1, 2}, {2, 3}, {3, 5}, {5, 8}, {8, 13}, {13, 21}, {21, 34}, {34, 55}, {55, 89}}
	if !cmp.Equal(got, want) {
		t.Errorf("Iterate = %+v, want %+v", got, want)
	}
}
