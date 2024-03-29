package enum

import (
	"golang.org/x/exp/constraints"
)

// Max returns the maximal element in the slice
// (or the zero value for an empty slice).
//
// Deprecated: As of Go 1.21, max is builtin: https://go.dev/ref/spec#Min_and_max
func Max[T constraints.Ordered](values []T) (ret T) {
	for i, v := range values {
		if i == 0 || v > ret {
			ret = v
		}
	}
	return ret
}

// Min returns the minimal element in the slice
// (or the zero value for an empty slice).
//
// Deprecated: As of Go 1.21, min is builtin: https://go.dev/ref/spec#Min_and_max
func Min[T constraints.Ordered](values []T) (ret T) {
	for i, v := range values {
		if i == 0 || v < ret {
			ret = v
		}
	}
	return ret
}

// MaxFunc returns the maximal element in the slice
// (or the zero value for an empty slice) using the
// provided lessFunc.
func MaxFunc[T any](values []T, lessFunc func(a, b T) bool) (ret T) {
	for i, v := range values {
		if i == 0 || !lessFunc(v, ret) {
			ret = v
		}
	}
	return ret
}

// MinFunc returns the minimal element in the slice
// (or the zero value for an empty slice) using the
// provided lessFunc.
func MinFunc[T any](values []T, lessFunc func(a, b T) bool) (ret T) {
	for i, v := range values {
		if i == 0 || lessFunc(v, ret) {
			ret = v
		}
	}
	return ret
}
