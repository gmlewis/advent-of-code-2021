package strfn

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	t.Parallel()
	if got, want := Length("yo"), 2; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
}

func TestSubstr(t *testing.T) {
	t.Parallel()
	f := Substr(2, 5)
	if got, want := f("yo howdy"), " ho"; got != want {
		t.Errorf("Substr = %v, want %v", got, want)
	}
}

func TestFirst(t *testing.T) {
	t.Parallel()
	if got, want := First("yo"), "y"; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}
}

func TestLast(t *testing.T) {
	t.Parallel()
	if got, want := Last("yo"), "o"; got != want {
		t.Errorf("Last = %v, want %v", got, want)
	}
}

func TestRunesWithIndex(t *testing.T) {
	t.Parallel()
	var got string
	RunesWithIndex("yo ho", func(i int, r rune) {
		got += fmt.Sprintf("[%v]%c;", i, r)
	})
	if want := "[0]y;[1]o;[2] ;[3]h;[4]o;"; got != want {
		t.Errorf("RunesWithIndex = %v, want %v", got, want)
	}
}
