package stream

import "golang.org/x/exp/slices"

// PermutationsOf streams all permutations of the given sequence to
// the returned channel.
func PermutationsOf[T any](sequence []T) <-chan []T {
	outCh := make(chan []T, defaultBufSize)

	go func() {
		permutationsOf(outCh, sequence)
		close(outCh)
	}()

	return outCh
}

func permutationsOf[T any](ch chan<- []T, seq []T) [][]T {
	if len(seq) == 0 {
		return nil
	}
	if len(seq) <= 1 {
		result := [][]T{seq}
		if ch != nil {
			ch <- seq
			return nil
		}
		return result
	}

	subperms := permutationsOf(nil, seq[1:])
	result := [][]T{}
	for _, sp := range subperms {
		for i := 0; i <= len(sp); i++ {
			subseq := slices.Insert(sp, i, seq[0])
			if ch != nil {
				ch <- subseq
				continue
			}
			result = append(result, subseq)
		}
	}
	return result
}
