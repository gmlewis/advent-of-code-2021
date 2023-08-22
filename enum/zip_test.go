package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestZip(t *testing.T) {
	t.Parallel()
	want := [][]int{{1, 3, 5}, {2, 4, 6}}
	got := Zip([][]int{{1, 2}, {3, 4}, {5, 6}})
	if !cmp.Equal(got, want) {
		t.Errorf("Zip = %+v, want %+v", got, want)
	}
}

func TestZip2(t *testing.T) {
	t.Parallel()
	type ns struct {
		N int
		S string
	}
	want := []ns{{1, "a"}, {2, "b"}, {3, "c"}}
	f := func(n int, s string) ns { return ns{n, s} }
	got := Zip2([]int{1, 2, 3, 4, 5, 6}, []string{"a", "b", "c"}, f)
	if !cmp.Equal(got, want) {
		t.Errorf("Zip2 = %+v, want %+v", got, want)
	}
}
