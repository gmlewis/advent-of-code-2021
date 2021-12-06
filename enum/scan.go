package enum

// Scan applies the given function to each element,
// emits the result and uses the same result as the accumulator for the
// next computation. It uses the given acc as the starting value.
//
// For example:
//   Scan(Range(1,5), 0, func(a, b int) int { return a + b })
// returns:
//   []int{1, 3, 6, 10, 15}
func Scan[T any](items []T, acc T, f func(a, b T) T) (ret []T) {
	for _, v := range items {
		nv := f(acc, v)
		ret = append(ret, nv)
		acc = nv
	}

	return ret
}
