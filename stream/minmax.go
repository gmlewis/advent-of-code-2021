package stream

import (
	"constraints"
)

// Max returns the maximal element in the channel
// (or the zero value for an empty channel).
func Max[T constraints.Ordered](ch <-chan T) (ret T) {
	var i int
	for v := range ch {
		if i == 0 || v > ret {
			ret = v
		}
		i++
	}
	return ret
}

// Min returns the minimal element in the channel
// (or the zero value for an empty channel).
func Min[T constraints.Ordered](ch <-chan T) (ret T) {
	var i int
	for v := range ch {
		if i == 0 || v < ret {
			ret = v
		}
		i++
	}
	return ret
}
