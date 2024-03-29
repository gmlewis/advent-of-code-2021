package algorithm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPriorityQueue_KeyOnly(t *testing.T) {
	t.Parallel()
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

func TestPriorityQueue_FixAndRemove(t *testing.T) {
	t.Parallel()
	type keyT [2]int
	m := map[keyT]int{
		{0, 0}: 1,
		{1, 0}: 2,
		{2, 0}: 3,
		{0, 1}: 4,
		{1, 1}: 5,
		{2, 1}: 6,
		{0, 2}: 7,
		{1, 2}: 8,
		{2, 2}: 9,
	}

	q := NewPriorityQueue(func(a, b keyT) bool { return m[a] < m[b] })

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			q.Push(keyT{x, y})
		}
	}

	// q.Remove(keyT{1, 1}) - panics!
	m[keyT{1, 2}] = -1
	q.Fix(keyT{1, 2})

	var got []keyT
	for i := 0; i < 9; i++ {
		got = append(got, q.Pop())
	}

	want := []keyT{
		{1, 2},
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{0, 2},
		{2, 2},
	}

	if !cmp.Equal(got, want) {
		t.Errorf("queue = %+v, want %+v", got, want)
	}
}
