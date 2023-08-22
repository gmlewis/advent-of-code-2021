package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFirst(t *testing.T) {
	t.Parallel()
	if got, want := First[int](nil), 0; got != want {
		t.Errorf("First(nil) = %v, want %v", got, want)
	}
	if got, want := First(ToChan([]int{1})), 1; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}
	if got, want := First(ToChan([]int{1, 2, 3})), 1; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}

	if got, want := First[string](nil), ""; got != want {
		t.Errorf("First(nil) = %v, want %v", got, want)
	}
	if got, want := First(ToChan([]string{"1"})), "1"; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}
	if got, want := First(ToChan([]string{"1", "2", "3"})), "1"; got != want {
		t.Errorf("First = %v, want %v", got, want)
	}
}

func TestLength(t *testing.T) {
	t.Parallel()
	if got, want := Length[int](nil), 0; got != want {
		t.Errorf("Length(nil) = %v, want %v", got, want)
	}
	if got, want := Length(ToChan([]int{1})), 1; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
	if got, want := Length(ToChan([]int{1, 2, 3})), 3; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}

	if got, want := Length[string](nil), 0; got != want {
		t.Errorf("Length(nil) = %v, want %v", got, want)
	}
	if got, want := Length(ToChan([]string{"1"})), 1; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
	if got, want := Length(ToChan([]string{"1", "2", "3"})), 3; got != want {
		t.Errorf("Length = %v, want %v", got, want)
	}
}

func TestTake(t *testing.T) {
	t.Parallel()
	if got, want := Take[int](nil, 5), []int(nil); !cmp.Equal(got, want) {
		t.Errorf("Take(nil) = %v, want %v", got, want)
	}
	if got, want := Take(ToChan([]int{1}), 5), []int{1}; !cmp.Equal(got, want) {
		t.Errorf("Take = %v, want %v", got, want)
	}
	if got, want := Take(ToChan([]int{1, 2, 3}), 2), []int{1, 2}; !cmp.Equal(got, want) {
		t.Errorf("Take = %v, want %v", got, want)
	}
	if got, want := Take(ToChan([]int{1, 2, 3}), 5), []int{1, 2, 3}; !cmp.Equal(got, want) {
		t.Errorf("Take = %v, want %v", got, want)
	}

	if got, want := Take[string](nil, 5), []string(nil); !cmp.Equal(got, want) {
		t.Errorf("Take(nil) = %v, want %v", got, want)
	}
	if got, want := Take(ToChan([]string{"1"}), 5), []string{"1"}; !cmp.Equal(got, want) {
		t.Errorf("Take = %v, want %v", got, want)
	}
	if got, want := Take(ToChan([]string{"1", "2", "3"}), 2), []string{"1", "2"}; !cmp.Equal(got, want) {
		t.Errorf("Take = %v, want %v", got, want)
	}
	if got, want := Take(ToChan([]string{"1", "2", "3"}), 5), []string{"1", "2", "3"}; !cmp.Equal(got, want) {
		t.Errorf("Take = %v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Parallel()
	if got, want := Sum[int](nil), 0; got != want {
		t.Errorf("Sum(nil) = %v, want %v", got, want)
	}
	if got, want := Sum(ToChan([]int{1})), 1; got != want {
		t.Errorf("Sum = %v, want %v", got, want)
	}
	if got, want := Sum(ToChan([]int{1, 2, 3, 4})), 10; got != want {
		t.Errorf("Sum = %v, want %v", got, want)
	}
}

func TestProduct(t *testing.T) {
	t.Parallel()
	if got, want := Product[int](nil), 0; got != want {
		t.Errorf("Product(nil) = %v, want %v", got, want)
	}
	if got, want := Product(ToChan([]int{1})), 1; got != want {
		t.Errorf("Product = %v, want %v", got, want)
	}
	if got, want := Product(ToChan([]int{1, 2, 3, 4})), 24; got != want {
		t.Errorf("Product = %v, want %v", got, want)
	}
}
