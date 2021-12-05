package strfn

import "testing"

func TestLength(t *testing.T) {
	if got, want := Length("yo"), 2; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
}
