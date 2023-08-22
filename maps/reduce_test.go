package maps

import (
	"testing"
)

func TestReduce(t *testing.T) {
	t.Parallel()
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	f := func(k string, v int, acc int) int { return acc + v }
	if got, want := Reduce(m, 0, f), 3; got != want {
		t.Errorf("Reduce = %v, want %v", got, want)
	}
}
