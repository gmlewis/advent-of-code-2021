package stream

import (
	"log"
)

// ChunkEvery takes a channel of values and chunks them n-at-a-time
// with the given step size. It discards any left-over items.
func ChunkEvery[T any](ch <-chan T, n, step int) <-chan []T {
	if step > n {
		log.Fatalf("step(%v) must be <= n(%v)", step, n)
	}

	outCh := make(chan []T, defaultBufSize)

	go func() {
		chunk := make([]T, 0, n)
		for v := range ch {
			chunk = append(chunk, v)
			if len(chunk) == n {
				outCh <- chunk
				chunk = chunk[step:]
			}
		}

		close(outCh)
	}()

	return outCh
}
