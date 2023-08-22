package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEquals(t *testing.T) {
	t.Parallel()
	eq1 := Equals(1)
	if got, want := eq1(1), true; got != want {
		t.Errorf("eq1(1) = %v, want %v", got, want)
	}
	if got, want := eq1(2), false; got != want {
		t.Errorf("eq1(2) = %v, want %v", got, want)
	}

	eqA := Equals("A")
	if got, want := eqA("A"), true; got != want {
		t.Errorf("eqA('A') = %v, want %v", got, want)
	}
	if got, want := eqA("2"), false; got != want {
		t.Errorf("eqA('2') = %v, want %v", got, want)
	}
}

func TestIdentity(t *testing.T) {
	t.Parallel()
	if got, want := Identity(1), 1; got != want {
		t.Errorf("Identity(1) = %v, want %v", got, want)
	}
	if got, want := Identity("A"), "A"; got != want {
		t.Errorf("Identity('A') = %v, want %v", got, want)
	}
}

func TestLength(t *testing.T) {
	t.Parallel()
	if got, want := Length[int](nil), 0; got != want {
		t.Errorf("Length(nil) = %v, want %v", got, want)
	}
	if got, want := Length([]int{1}), 1; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
	if got, want := Length([]int{1, 2, 3}), 3; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}

	if got, want := Length[string](nil), 0; got != want {
		t.Errorf("Length(nil) = %v, want %v", got, want)
	}
	if got, want := Length([]string{"1"}), 1; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
	if got, want := Length([]string{"1", "2", "3"}), 3; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
}

func TestFirst(t *testing.T) {
	t.Parallel()
	if got, want := First[int](nil), 0; got != want {
		t.Errorf("First(nil) = %v, want %v", got, want)
	}
	if got, want := First([]int{1}), 1; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}
	if got, want := First([]int{1, 2, 3}), 1; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}

	if got, want := First[string](nil), ""; got != want {
		t.Errorf("First(nil) = %v, want %v", got, want)
	}
	if got, want := First([]string{"1"}), "1"; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}
	if got, want := First([]string{"1", "2", "3"}), "1"; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}
}

func TestLonger(t *testing.T) {
	t.Parallel()
	if got, want := Longer[int](nil, nil), []int(nil); !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer(nil, []int{1}), []int{1}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer([]int{1}, nil), []int{1}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer([]int{1}, []int{1, 2, 3}), []int{1, 2, 3}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer([]int{1, 2, 3}, []int{1}), []int{1, 2, 3}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}

	if got, want := Longer[string](nil, nil), []string(nil); !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer(nil, []string{"1"}), []string{"1"}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer([]string{"1"}, nil), []string{"1"}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer([]string{"1"}, []string{"1", "2", "3"}), []string{"1", "2", "3"}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
	if got, want := Longer([]string{"1", "2", "3"}, []string{"1"}), []string{"1", "2", "3"}; !cmp.Equal(got, want) {
		t.Errorf("Longer = %v, want %v", got, want)
	}
}

func TestShorter(t *testing.T) {
	t.Parallel()
	if got, want := Shorter[int](nil, nil), []int(nil); !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter(nil, []int{1}), []int(nil); !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter([]int{1}, nil), []int(nil); !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter([]int{1}, []int{1, 2, 3}), []int{1}; !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter([]int{1, 2, 3}, []int{1}), []int{1}; !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}

	if got, want := Shorter[string](nil, nil), []string(nil); !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter(nil, []string{"1"}), []string(nil); !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter([]string{"1"}, nil), []string(nil); !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter([]string{"1"}, []string{"1", "2", "3"}), []string{"1"}; !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
	if got, want := Shorter([]string{"1", "2", "3"}, []string{"1"}), []string{"1"}; !cmp.Equal(got, want) {
		t.Errorf("Shorter = %v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Parallel()
	if got, want := Sum[int](nil), 0; got != want {
		t.Errorf("Sum(nil) = %v, want %v", got, want)
	}
	if got, want := Sum([]int{1}), 1; got != want {
		t.Errorf("Sum = %v, want %v", got, want)
	}
	if got, want := Sum([]int{1, 2, 3, 4}), 10; got != want {
		t.Errorf("Sum = %v, want %v", got, want)
	}
}

func TestProduct(t *testing.T) {
	t.Parallel()
	if got, want := Product[int](nil), 0; got != want {
		t.Errorf("Product(nil) = %v, want %v", got, want)
	}
	if got, want := Product([]int{1}), 1; got != want {
		t.Errorf("Product = %v, want %v", got, want)
	}
	if got, want := Product([]int{1, 2, 3, 4}), 24; got != want {
		t.Errorf("Product = %v, want %v", got, want)
	}
}
