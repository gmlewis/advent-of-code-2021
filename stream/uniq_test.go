package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniq(t *testing.T) {
	want := []int{1, 2, 3}
	got := ToSlice(Uniq(ToChan([]int{1, 2, 3, 3, 2, 1})))
	if !cmp.Equal(got, want) {
		t.Errorf("Uniq = %+v, want %+v", got, want)
	}
}
