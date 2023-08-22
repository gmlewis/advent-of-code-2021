package enum

import (
	"log"
	"math"
)

// Range returns a slice of numbers that run from start to end.
func Range[T Number](start, end T) (ret []T) {
	d := T(1)
	n := int(1 + T(end-start))
	if start > end {
		d = -d
		n = int(1 + T(start-end))
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
// Note that one of the ranges might overshoot if the
// distances are not identical.
func Ranges[T int](start, end []T) (ret [][]T) {
	if len(start) != len(end) {
		log.Fatal("start and end must be same length")
	}

	inc := make([]float64, len(start))
	steps := 1
	for i, v := range start {
		switch {
		case end[i]-v > T(0):
			inc[i] = 1
		case end[i]-v < T(0):
			inc[i] = -1
		}

		s := 1 + int(0.5+inc[i]*float64(end[i]-v))
		if s > steps {
			steps = s
		}
	}

	for i, v := range start {
		inc[i] = inc[i] * (1 + math.Abs(float64(end[i]-v))) / float64(steps)
	}

	for i := 0; i < steps; i++ {
		v := make([]T, len(start))
		for j, d := range inc {
			v[j] = T(math.Round(float64(start[j]) + float64(i)*d))
		}
		ret = append(ret, v)
	}

	return ret
}
