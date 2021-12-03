package must

// ChunkEvery takes a slice of values and chunks them n-at-a-time
// with the given step size. It discards any left-over items.
func ChunkEvery[T any](vals []T, n, step int) [][]T {
	result := [][]T{}
	for i := 0; i+n-1 < len(vals); i += step {
		chunk := make([]T, 0, n)
		for j := 0; j < n; j++ {
			chunk = append(chunk, vals[i+j])
		}
		result = append(result, chunk)
	}
	return result
}
