package must

import "testing"

func TestAtoi(t *testing.T) {
	if got, want := Atoi("333"), 333; got != want {
		t.Errorf("Atoi = %v, want %v", got, want)
	}
}
