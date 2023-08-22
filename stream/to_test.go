package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToChan_ToSlice(t *testing.T) {
	t.Parallel()
	want := []int{1, 2, 3}
	got := ToSlice(ToChan(want))
	if !cmp.Equal(got, want) {
		t.Errorf("ToChan |> ToSlice = %+v, want %+v", got, want)
	}
}
