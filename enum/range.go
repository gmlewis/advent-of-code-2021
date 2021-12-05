package enum

import (
	"log"
)

// Range returns a slice of numbers that run from start to end.
//
// For example:
//   Ranges(0, 2), Ranges(2, 0)
// respectively return:
//   []int{0, 1, 2}, []int{2, 1, 0}
func Range[T int](start, end T) (ret []T) {
	d := T(1)
	n := 1 + int(end-start)
	if start > end {
		d = T(-1)
		n = 1 + int(start-end)
	}
	ret = make([]T, 0, n)

	last := start
	for i := 0; i < n; i++ {
		ret = append(ret, last)
		last += d
	}

	return ret
}

// Ranges returns a slice of slices-of-integers that
// increment all values from the start to the end.
//
// For example:
//   Ranges([]int{0,3,0}, []int{2,1,0})
// returns:
//   [][]int{{0,3,0}, {1,2,0}, {2,1,0}}
//
// Note that one of the ranges might overshoot if the
// distances are not identical.
func Ranges[T int](start, end []T) (ret [][]T) {
	if len(start) != len(end) {
		log.Fatal("start and end must be same length")
	}

	inc := make([]T, len(start))
	steps := 1
	for i, v := range start {
		switch {
		case end[i]-v > T(0):
			inc[i] = T(1)
		case end[i]-v < T(0):
			inc[i] = T(-1)
		}

		s := int(1 + inc[i]*(end[i]-v))
		if s > steps {
			steps = s
		}
	}

	last := make([]T, len(start))
	copy(last, start)
	for i := 0; i < steps; i++ {
		v := make([]T, len(start))
		copy(v, last)
		ret = append(ret, v)
		for j, d := range inc {
			last[j] += d
		}
	}

	return ret
}
