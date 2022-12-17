package stream

import "golang.org/x/exp/slices"

// PermutationsOf streams all permutations of the given sequence to
// the returned channel.
func PermutationsOf[T any](sequence []T) <-chan []T {
	outCh := make(chan []T, 5)

	go permutationsOf(outCh, sequence)

	return outCh
}

func permutationsOf[T any](ch chan<- []T, seq []T) {
	defer close(ch)

	if len(seq) == 0 {
		return
	}
	if len(seq) <= 1 {
		ch <- seq
		return
	}

	childCh := make(chan []T, 5)
	go permutationsOf(childCh, seq[1:])

	for sp := range childCh {
		for i := 0; i <= len(sp); i++ {
			ch <- slices.Insert(sp, i, seq[0])
		}
	}
}
