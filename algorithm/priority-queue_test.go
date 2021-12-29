package algorithm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPriorityQueue(t *testing.T) {
	type keyT [2]int
	less := func(a, b keyT) bool {
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] > b[1]
	}
	q := NewPriorityQueue(less)

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			q.Push(keyT{x, y})
		}
	}

	var got []keyT
	for i := 0; i < 9; i++ {
		got = append(got, q.Pop())
	}

	want := []keyT{
		{0, 2},
		{1, 2},
		{2, 2},
		{0, 1},
		{1, 1},
		{2, 1},
		{0, 0},
		{1, 0},
		{2, 0},
	}

	if !cmp.Equal(got, want) {
		t.Errorf("queue = %+v, want %+v", got, want)
	}
}
