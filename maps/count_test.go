package maps

import (
	"testing"
)

func TestCount(t *testing.T) {
	t.Parallel()
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	f := func(k string, v int) bool { return v > 0 }
	if got, want := Count(m, f), 2; got != want {
		t.Errorf("Count = %v, want %v", got, want)
	}
}
