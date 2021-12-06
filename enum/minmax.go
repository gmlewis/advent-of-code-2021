package enum

import "constraints"

// Max returns the maximal element in the slice
// (or the zero value for an empty slice).
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
func Min[T constraints.Ordered](values []T) (ret T) {
	for i, v := range values {
		if i == 0 || v < ret {
			ret = v
		}
	}
	return ret
}
