package tuple

import "testing"

func TestNew3(t *testing.T) {
	t3 := New3(0, "1", uint(2))
	if got, want := t3.A, 0; got != want {
		t.Errorf("New3.A = %v, want %v", got, want)
	}
	if got, want := t3.B, "1"; got != want {
		t.Errorf("New3.B = %v, want %v", got, want)
	}
	if got, want := t3.C, uint(2); got != want {
		t.Errorf("New3.C = %v, want %v", got, want)
	}
}
