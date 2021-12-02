package must

// ChunkEveryInt takes a slice of ints and chunks them n-at-a-time
// with the given step size. It discards any left-over items.
func ChunkEveryInt(vals []int, n, step int) [][]int {
	result := [][]int{}
	for i := 0; i+n-1 < len(vals); i += step {
		chunk := make([]int, 0, n)
		for j := 0; j < n; j++ {
			chunk = append(chunk, vals[i+j])
		}
		result = append(result, chunk)
	}
	return result
}
