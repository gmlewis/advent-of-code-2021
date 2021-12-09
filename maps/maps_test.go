package maps

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHasKey(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	f := HasKey(m)
	if !f("a") {
		t.Error("HasKey = false, want true")
	}
	if f("yo") {
		t.Error("HasKey = true, want false")
	}
}

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

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	got := Keys(m)
	sort.Strings(got)
	if want := []string{"a", "b", "c"}; !cmp.Equal(got, want) {
		t.Errorf("Keys = %v, want %v", got, want)
	}
	if got, want := Keys(map[int]int{}), []int{}; !cmp.Equal(got, want) {
		t.Errorf("Keys = %v, want %v", got, want)
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

func TestSwap(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	got := Swap(m)
	want := map[int]string{0: "a", 1: "b", 2: "c"}
	if !cmp.Equal(got, want) {
		t.Errorf("Swap = %v, want %v", got, want)
	}
}

func TestSumKeys(t *testing.T) {
	m := map[int]int{7: 0, 8: 1, 9: 2}
	if got, want := SumKeys(m), 24; got != want {
		t.Errorf("SumKeys = %v, want %v", got, want)
	}
}

func TestSumValues(t *testing.T) {
	m := map[int]int{7: 0, 8: 1, 9: 2}
	if got, want := SumValues(m), 3; got != want {
		t.Errorf("SumValues = %v, want %v", got, want)
	}
}

func TestProductKeys(t *testing.T) {
	m := map[int]int{7: 0, 8: 1, 9: 2}
	if got, want := ProductKeys(m), 504; got != want {
		t.Errorf("ProductKeys = %v, want %v", got, want)
	}
}

func TestProductValues(t *testing.T) {
	m := map[int]int{7: 3, 8: 1, 9: 2}
	if got, want := ProductValues(m), 6; got != want {
		t.Errorf("ProductValues = %v, want %v", got, want)
	}
}
