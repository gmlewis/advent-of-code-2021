package stream

import (
	"log"
	"math"
)

// Range returns a channel of numbers that run from start to end.
//
// For example:
//   Ranges(0, 2), Ranges(2, 0)
// respectively return:
//   []int{0, 1, 2}, []int{2, 1, 0}
func Range[T int](start, end T) <-chan T {
	d := T(1)
	n := 1 + int(end-start)
	if start > end {
		d = T(-1)
		n = 1 + int(start-end)
	}
	size := defaultBufSize
	if n < size {
		size = n
	}
	ch := make(chan T, size)

	go func() {
		last := start
		for i := 0; i < n; i++ {
			ch <- last
			last += d
		}
		close(ch)
	}()

	return ch
}

// Ranges returns a channel of slices-of-integers that
// increment all values from the start to the end.
//
// For example:
//   Ranges([]int{0,3,0}, []int{2,1,0})
// returns:
//   [][]int{{0,3,0}, {1,2,0}, {2,1,0}}
//
// Note that one of the ranges might overshoot if the
// distances are not identical.
func Ranges[T int](start, end []T) <-chan []T {
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

	size := defaultBufSize
	if steps < size {
		size = steps
	}
	ch := make(chan []T, size)

	go func() {
		for i := 0; i < steps; i++ {
			v := make([]T, len(start))
			for j, d := range inc {
				v[j] = T(math.Round(float64(start[j]) + float64(i)*d))
			}
			ch <- v
		}

		close(ch)
	}()

	return ch
}
