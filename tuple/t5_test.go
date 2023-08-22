package tuple

import "testing"

func TestNew5(t *testing.T) {
	t.Parallel()
	t5 := New5(0, "1", uint(2), 3.0, 4+0i)
	if got, want := t5.A, 0; got != want {
		t.Errorf("New5.A = %v, want %v", got, want)
	}
	if got, want := t5.B, "1"; got != want {
		t.Errorf("New5.B = %v, want %v", got, want)
	}
	if got, want := t5.C, uint(2); got != want {
		t.Errorf("New5.C = %v, want %v", got, want)
	}
	if got, want := t5.D, 3.0; got != want {
		t.Errorf("New5.D = %v, want %v", got, want)
	}
	if got, want := t5.E, 4+0i; got != want {
		t.Errorf("New5.E = %v, want %v", got, want)
	}
}
