package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScan(t *testing.T) {
	want := []int{1, 3, 6, 10, 15}
	got := Scan(Range(1, 5), 0, func(a, b int) int { return a + b })
	if !cmp.Equal(got, want) {
		t.Errorf("Scan = %+v, want %+v", got, want)
	}
}
