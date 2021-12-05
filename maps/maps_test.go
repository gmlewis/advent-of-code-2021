package maps

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKey(t *testing.T) {
	if got, want := Key(0, "value"), 0; got != want {
		t.Errorf("Key = %v, want %v", got, want)
	}

	if got, want := Key("key", "value"), "key"; got != want {
		t.Errorf("Key = %v, want %v", got, want)
	}
}

func TestValue(t *testing.T) {
	if got, want := Value(0, "value"), "value"; got != want {
		t.Errorf("Value = %v, want %v", got, want)
	}

	if got, want := Value("key", 0), 0; got != want {
		t.Errorf("Value = %v, want %v", got, want)
	}

	if got, want := Value("key", []int{1, 2, 3}), []int{1, 2, 3}; !cmp.Equal(got, want) {
		t.Errorf("Value = %v, want %v", got, want)
	}
}

func TestValueLen(t *testing.T) {
	if got, want := ValueLen(0, []int{1, 2, 3}), 3; got != want {
		t.Errorf("ValueLen = %v, want %v", got, want)
	}

	if got, want := ValueLen("key", []string{"a", "b"}), 2; got != want {
		t.Errorf("ValueLen = %v, want %v", got, want)
	}

	if got, want := ValueLen("key", []int{}), 0; got != want {
		t.Errorf("ValueLen = %v, want %v", got, want)
	}

	if got, want := ValueLen[string, int]("key", nil), 0; got != want {
		t.Errorf("ValueLen = %v, want %v", got, want)
	}
}

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

func TestAny(t *testing.T) {
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

func TestMap(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	f := func(k string, v int) string { return fmt.Sprintf("%v:%v", k, v) }
	got := Map(m, f)
	want := []string{"a:0", "b:1", "c:2"}
	sort.Strings(got)
	if !cmp.Equal(got, want) {
		t.Errorf("Map = %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	f := func(k string, v int, acc int) int { return acc + v }
	if got, want := Reduce(m, 0, f), 3; got != want {
		t.Errorf("Reduce = %v, want %v", got, want)
	}
}
