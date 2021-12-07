package tuple

import "testing"

func TestNew4(t *testing.T) {
	t4 := New4(0, "1", uint(2), 3.0)
	if got, want := t4.A, 0; got != want {
		t.Errorf("New4.A = %v, want %v", got, want)
	}
	if got, want := t4.B, "1"; got != want {
		t.Errorf("New4.B = %v, want %v", got, want)
	}
	if got, want := t4.C, uint(2); got != want {
		t.Errorf("New4.C = %v, want %v", got, want)
	}
	if got, want := t4.D, 3.0; got != want {
		t.Errorf("New4.D = %v, want %v", got, want)
	}
}
