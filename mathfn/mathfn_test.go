package mathfn

import "testing"

func TestAbs(t *testing.T) {
	if got, want := Abs(-1), 1; got != want {
		t.Errorf("Abs = %v, want %v", got, want)
	}
	if got, want := Abs(1), 1; got != want {
		t.Errorf("Abs = %v, want %v", got, want)
	}
	if got, want := Abs(-1.0), 1.0; got != want {
		t.Errorf("Abs = %v, want %v", got, want)
	}
	if got, want := Abs(1.0), 1.0; got != want {
		t.Errorf("Abs = %v, want %v", got, want)
	}
}
