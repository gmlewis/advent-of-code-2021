package maps

import (
	"testing"
)

func TestAll(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	f1 := func(k string, v int) bool { return v < 10 }
	if got, want := All(m, f1), true; got != want {
		t.Errorf("All = %v, want %v", got, want)
	}

	f2 := func(k string, v int) bool { return v < 2 }
	if got, want := All(m, f2), false; got != want {
		t.Errorf("All = %v, want %v", got, want)
	}
}
