package enum

import "testing"

func TestReduce(t *testing.T) {
	t.Parallel()
	f := func(v, acc int) int { return acc + v*v }
	got := Reduce([]int{1, 2, 3}, 0, f)
	if want := 14; got != want {
		t.Errorf("Reduce = %+v, want %+v", got, want)
	}
}

func TestReduceWithIndex(t *testing.T) {
	t.Parallel()
	f := func(i, v, acc int) int { return acc + v*v + i*i*i }
	got := ReduceWithIndex([]int{1, 2, 3}, 0, f)
	if want := 23; got != want {
		t.Errorf("Reduce = %+v, want %+v", got, want)
	}
}
