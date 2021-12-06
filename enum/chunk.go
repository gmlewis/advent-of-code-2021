package enum

// ChunkEvery takes a slice of items and chunks them n-at-a-time
// with the given step size. It discards any left-over items.
func ChunkEvery[T any](items []T, n, step int) [][]T {
	result := [][]T{}
	if n <= 0 || step <= 0 {
		return result
	}

	for i := 0; i+n-1 < len(items); i += step {
		chunk := make([]T, 0, n)
		for j := 0; j < n; j++ {
			chunk = append(chunk, items[i+j])
		}
		result = append(result, chunk)
	}
	return result
}
