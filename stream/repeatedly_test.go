package stream

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepeatedly(t *testing.T) {
	rand.Seed(0)
	want := []int32{2029793274, 526058514, 1408655353, 116702506, 789387515,
		621654496, 413258767, 1407315077, 1926657288, 359390928}
	got := Take(Repeatedly(rand.Int31), 10)
	if !cmp.Equal(got, want) {
		t.Errorf("Repeatedly = %+v, want %+v", got, want)
	}
}
