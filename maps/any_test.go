package maps

import (
	"testing"
)

func TestAny(t *testing.T) {
	t.Parallel()
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	f1 := func(k string, v int) bool { return v < 10 }
	if got, want := Any(m, f1), true; got != want {
		t.Errorf("Any = %v, want %v", got, want)
	}

	f2 := func(k string, v int) bool { return v < 0 }
	if got, want := Any(m, f2), false; got != want {
		t.Errorf("Any = %v, want %v", got, want)
	}
}
