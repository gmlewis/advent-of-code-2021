package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestZip2(t *testing.T) {
	type ns struct {
		N int
		S string
	}
	want := []ns{{1, "a"}, {2, "b"}, {3, "c"}}
	f := func(n int, s string) ns { return ns{n, s} }
	got := ToSlice(Zip2(ToChan([]int{1, 2, 3, 4, 5, 6}), ToChan([]string{"a", "b", "c"}), f))
	if !cmp.Equal(got, want) {
		t.Errorf("Zip2 = %+v, want %+v", got, want)
	}
}
