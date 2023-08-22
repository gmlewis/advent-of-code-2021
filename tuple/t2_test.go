package tuple

import "testing"

func TestNew2(t *testing.T) {
	t.Parallel()
	t2 := New2(0, "1")
	if got, want := t2.A, 0; got != want {
		t.Errorf("New2.A = %v, want %v", got, want)
	}
	if got, want := t2.B, "1"; got != want {
		t.Errorf("New2.B = %v, want %v", got, want)
	}
}
